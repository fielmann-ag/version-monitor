package release

import (
	"errors"
)

// Adapter Error
var (
	ErrOwnerEmpty = errors.New("gitHubRelease.owner must not be empty")
	ErrRepoEmpty = errors.New("gitHubRelease.repo must not be empty")
)
