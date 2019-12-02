package monitor

import (
	"time"
)

// An Adapter is used to load a version string from a specific target / technology / transport
type Adapter interface {
	Fetch(cfg AdapterConfig) (string, error)
	Validate(cfg AdapterConfig) error
}

// AdapterType defines the adapter to use to load the version from
type AdapterType string

// AdapterRegistry is a registry for adapters ... wow
type AdapterRegistry map[AdapterType]Adapter

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
