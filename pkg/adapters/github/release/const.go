package release

import (
	"errors"

	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// Adapter Error
var (
	ErrOwnerEmpty = errors.New("gitHubRelease.owner must not be empty")
	ErrRepoEmpty  = errors.New("gitHubRelease.repo must not be empty")
)

const (
	// AdapterType constant
	AdapterType monitor.AdapterType = "gitHubRelease"
)
