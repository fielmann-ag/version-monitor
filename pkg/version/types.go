package version

import (
	"time"

	"github.com/fielmann-ag/ops-version-monitor/pkg/config"
)

// An Adapter is used to load a version string from a specific target / technology / transport
type Adapter interface {
	Fetch(cfg config.AdapterConfig) (string, error)
	Validate(cfg config.AdapterConfig) error
}

// Version information about a specific system / technology
type Version struct {
	Name    string
	Current string
	Latest  string
}

// A Monitor receives the latest set of versions fetched from adapters
type Monitor interface {
	Versions() ([]Version, time.Time, error)
}
