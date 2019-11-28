package periodic

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/internal/version"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// Monitor periodically iterates the map of adapters and updates the fetched versions
type Monitor struct {
	sync.RWMutex
	logger           logging.Logger
	config           *monitor.Config
	cachedVersions   map[string]monitor.Version
	lastError        error
	latestResultFrom time.Time
	adapters         monitor.AdapterRegistry
}

// NewMonitor returns a new fetcher instance
func NewMonitor(logger logging.Logger, config *monitor.Config, adapters monitor.AdapterRegistry) *Monitor {
	return &Monitor{
		logger:         logger,
		config:         config,
		cachedVersions: map[string]monitor.Version{},
		adapters:       adapters,
	}
}

// Versions returns the latest set of versions cached since the last update
func (m *Monitor) Versions() ([]monitor.Version, time.Time, error) {
	m.RLock()
	if m.lastError != nil {
		m.RUnlock()
		return nil, time.Time{}, m.lastError
	}

	versions := make([]monitor.Version, 0)
	for _, v := range m.cachedVersions {
		versions = append(versions, v)
	}

	m.RUnlock()
	return versions, m.latestResultFrom, nil
}

func (m *Monitor) validateConfig() error {
	if m.config == nil {
		return errors.New("config is not set")
	}

	if len(m.config.Targets) == 0 {
		return errors.New("no targets defined")
	}

	for _, t := range m.config.Targets {
		if _, ok := m.adapters[t.Latest.Type]; !ok {
			return fmt.Errorf("target.latest.type %q of target %s not found", t.Latest.Type, t.Name)
		}
		if _, ok := m.adapters[t.Current.Type]; !ok {
			return fmt.Errorf("target.current.type %q of target %s not found", t.Current.Type, t.Name)
		}
	}

	return nil
}

// Start the periodic fetching
func (m *Monitor) Start() error {
	if err := m.validateConfig(); err != nil {
		return err
	}

	c := cron.New()
	if _, err := c.AddJob("@hourly", m); err != nil {
		return fmt.Errorf("failed to register cron job: %v", err)
	}

	c.Start()
	m.Run()
	return nil
}

// Run fetch for all targets
func (m *Monitor) Run() {
	m.logger.Debugf("fetching versions ...")

	for _, target := range m.config.Targets {
		go m.fetch(target)
	}
}

func (m *Monitor) fetch(target monitor.Target) {
	m.logger.Debugf("fetching version %v", target.Name)

	currentVersionAdapter := m.adapters[target.Current.Type]
	currentVersion, err := currentVersionAdapter.Fetch(target.Current)
	if err != nil {
		m.error(fmt.Errorf("failed to load version from target.Current adapter %v: %v", target.Current.Type, err))
		return
	}

	latestVersionAdapter := m.adapters[target.Latest.Type]
	latestVersion, err := latestVersionAdapter.Fetch(target.Latest)
	if err != nil {
		m.error(fmt.Errorf("failed to load version from target.Latest adapter %v: %v", target.Latest.Type, err))
		return
	}

	m.storeVersion(target.Name, monitor.Version{
		Name:    target.Name,
		Current: version.Clean(currentVersion),
		Latest:  version.Clean(latestVersion),
	})

	m.logger.Debugf("fetching version %v done", target.Name)
}

func (m *Monitor) storeVersion(targetName string, version monitor.Version) {
	m.Lock()
	m.cachedVersions[targetName] = version
	m.latestResultFrom = time.Now()
	m.Unlock()
}

func (m *Monitor) error(err error) {
	m.logger.Error(err)
	m.Lock()
	m.lastError = err
	m.Unlock()
}
