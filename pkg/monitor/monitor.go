package monitor

import (
	"fmt"
	"sync"
	"time"

	"github.com/fielmann-ag/version-monitor/pkg/adapters"
	"github.com/fielmann-ag/version-monitor/pkg/config"
	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/version"

	"github.com/robfig/cron/v3"
)

// PeriodicMonitor periodically iterates the map of adapters and updates the fetched versions
type PeriodicMonitor struct {
	sync.RWMutex
	logger           logging.Logger
	config           *config.Config
	cachedVersions   map[string]version.Version
	lastError        error
	latestResultFrom time.Time
}

// NewPeriodic returns a new fetcher instance
func NewPeriodic(logger logging.Logger, config *config.Config) *PeriodicMonitor {
	return &PeriodicMonitor{
		logger:         logger,
		config:         config,
		cachedVersions: map[string]version.Version{},
	}
}

// Versions returns the latest set of versions cached since the last update
func (m *PeriodicMonitor) Versions() ([]version.Version, time.Time, error) {
	m.RLock()
	if m.lastError != nil {
		m.RUnlock()
		return nil, time.Time{}, m.lastError
	}

	versions := make([]version.Version, 0)
	for _, v := range m.cachedVersions {
		versions = append(versions, v)
	}

	m.RUnlock()
	return versions, m.latestResultFrom, nil
}

func (m *PeriodicMonitor) validateConfig() error {
	for _, t := range m.config.Targets {
		if _, ok := adapters.Registry[t.Latest.Type]; !ok {
			return fmt.Errorf("target.latest.type %q of target %s not found", t.Latest.Type, t.Name)
		}
		if _, ok := adapters.Registry[t.Current.Type]; !ok {
			return fmt.Errorf("target.current.type %q of target %s not found", t.Current.Type, t.Name)
		}
	}

	return nil
}

// Start the periodic fetching
func (m *PeriodicMonitor) Start() error {
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
func (m *PeriodicMonitor) Run() {
	m.logger.Debugf("fetching versions ...")

	for _, target := range m.config.Targets {
		go m.fetch(target)
	}
}

func (m *PeriodicMonitor) fetch(target config.Target) {
	m.logger.Debugf("fetching version %v", target.Name)

	currentVersionAdapter := adapters.Registry[target.Current.Type]
	currentVersion, err := currentVersionAdapter.Fetch(target.Current)
	if err != nil {
		m.error(fmt.Errorf("failed to load version from target.Current adapter %v: %v", target.Current.Type, err))
		return
	}

	latestVersionAdapter := adapters.Registry[target.Latest.Type]
	latestVersion, err := latestVersionAdapter.Fetch(target.Latest)
	if err != nil {
		m.error(fmt.Errorf("failed to load version from target.Latest adapter %v: %v", target.Latest.Type, err))
		return
	}

	m.storeVersion(target.Name, version.Version{
		Name:    target.Name,
		Current: cleanVersion(currentVersion),
		Latest:  cleanVersion(latestVersion),
	})

	m.logger.Debugf("fetching version %v done", target.Name)
}

func (m *PeriodicMonitor) storeVersion(targetName string, version version.Version) {
	m.Lock()
	m.cachedVersions[targetName] = version
	m.latestResultFrom = time.Now()
	m.Unlock()
}

func (m *PeriodicMonitor) error(err error) {
	m.logger.Error(err)
	m.Lock()
	m.lastError = err
	m.Unlock()
}
