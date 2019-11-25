package version

import (
	"fmt"
	"sync"
	"time"

	"github.com/fielmann-ag/ops-version-monitor/pkg/internal/logging"

	"github.com/robfig/cron/v3"
)

// PeriodicMonitor periodically iterates the map of adapters and updates the fetched versions
type PeriodicMonitor struct {
	sync.RWMutex
	logger           logging.Logger
	cachedVersions   map[string]Version
	lastError        error
	latestResultFrom time.Time
}

// NewPeriodicMonitor returns a new fetcher instance
func NewPeriodicMonitor(logger logging.Logger) *PeriodicMonitor {
	return &PeriodicMonitor{
		logger:         logger,
		cachedVersions: map[string]Version{},
	}
}

// Versions returns the latest set of versions cached since the last update
func (m *PeriodicMonitor) Versions() ([]Version, time.Time, error) {
	m.RLock()
	if m.lastError != nil {
		m.RUnlock()
		return nil, time.Time{}, m.lastError
	}

	versions := make([]Version, 0)
	for _, version := range m.cachedVersions {
		versions = append(versions, version)
	}

	m.RUnlock()
	return versions, m.latestResultFrom, nil
}

// Start the periodic fetching
func (m *PeriodicMonitor) Start() error {
	m.Run()

	c := cron.New()
	if _, err := c.AddJob("@hourly", m); err != nil {
		return fmt.Errorf("failed to register cron job: %v", err)
	}

	c.Start()
	return nil
}

// Run fetch from all adapters
func (m *PeriodicMonitor) Run() {
	m.logger.Debugf("start fetching versions ...")

	for name, adapter := range adapters {
		go m.fetchAdapter(name, adapter)
	}

	m.logger.Debugf("done fetching versions.")
}

func (m *PeriodicMonitor) fetchAdapter(name string, adapter Adapter) {
	m.logger.Debugf("fetching version from %v adapter", name)

	v, err := adapter.FetchVersion()
	if err != nil {
		m.error(name, err)
		return
	}

	m.storeVersion(name, v)
	m.logger.Debugf("done fetching version from %v adapter", name)
}

func (m *PeriodicMonitor) storeVersion(name string, version Version) {
	m.Lock()
	m.cachedVersions[name] = version
	m.Unlock()
}

func (m *PeriodicMonitor) error(name string, err error) {
	m.logger.Errorf("failed to fetch version from adapter %v: %v", name, err)
	m.Lock()
	m.lastError = err
	m.Unlock()
}
