package release

import (
	"errors"
)

// Adapter Error
var (
	ErrReleaseEmpty = errors.New("gitHubRelease.repo must not be empty")
)
