package testing

import (
	"time"

	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

var _ monitor.Monitor = &Monitor{}

// Monitor is a static mock for monitor.Monitor interface
type Monitor struct {
	versions []monitor.Version
	t        time.Time
	err      error
}

// NewMonitor returns a new Monitor instance
func NewMonitor(versions []monitor.Version, t time.Time, err error) *Monitor {
	return &Monitor{
		versions: versions,
		t:        t,
		err:      err,
	}
}

// Versions returns the versions set via constructor
func (m *Monitor) Versions() ([]monitor.Version, time.Time, error) {
	return m.versions, m.t, m.err
}
