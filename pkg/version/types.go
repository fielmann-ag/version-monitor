package version

import (
	"time"
)

// An Adapter is used to load a version string from a specific target / technology / transport
type Adapter interface {
	FetchVersion() (Version, error)
}

// Version information about a specific system / technology
type Version struct {
	Current string
	Latest  string
}

// A Monitor receives the latest set of versions fetched from adapters
type Monitor interface {
	Versions() ([]Version, time.Time, error)
}
