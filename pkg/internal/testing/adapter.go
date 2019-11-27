package testing

import (
	"github.com/fielmann-ag/ops-version-monitor/pkg/config"
)

// AdapterTypeStatic is the AdapterType value for the static adapter
const AdapterTypeStatic config.AdapterType = "static"

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

func (a *StaticAdapter) Validate(cfg config.AdapterConfig) error {
	return nil
}

// Fetch returns the static fields value, mocking the version.Adapter interface
func (a *StaticAdapter) Fetch(cfg config.AdapterConfig) (string, error) {
	return a.Version, nil
}
