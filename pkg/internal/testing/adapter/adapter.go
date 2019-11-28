package adapter

import (
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// AdapterTypeStatic is the AdapterType value for the static adapter
const AdapterTypeStatic monitor.AdapterType = "static"

// StaticAdapter is a simple test adapter returning a static value
type StaticAdapter struct {
	Version string
}

// NewStaticAdapter returns a new StaticAdapter instance with given value set
func NewStaticAdapter(version string) *StaticAdapter {
	return &StaticAdapter{
		Version: version,
	}
}

// Validate the given configuration
func (a *StaticAdapter) Validate(cfg monitor.AdapterConfig) error {
	return nil
}

// Fetch returns the static fields value, mocking the version.Adapter interface
func (a *StaticAdapter) Fetch(cfg monitor.AdapterConfig) (string, error) {
	return a.Version, nil
}
