package testing

import (
	"github.com/fielmann-ag/ops-version-monitor/pkg/version"
)

// StaticAdapter is a simple test adapter returning a static value
type StaticAdapter struct {
	Current, Latest string
}

// NewStaticAdapter returns a new StaticAdapter instance with given value set
func NewStaticAdapter(current, latest string) *StaticAdapter {
	return &StaticAdapter{
		Current: current,
		Latest:  latest,
	}
}

// FetchVersion returns the static fields value, mocking the version.Adapter interface
func (a *StaticAdapter) FetchVersion() (version.Version, error) {
	return version.Version{
		Current: a.Current,
		Latest:  a.Latest,
	}, nil
}
