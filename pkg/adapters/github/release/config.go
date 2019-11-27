package release

import (
	"github.com/google/go-github/v28/github"

	"github.com/fielmann-ag/ops-version-monitor/pkg/config"
	"github.com/fielmann-ag/ops-version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/ops-version-monitor/pkg/version"
)

// AdapterConstructor creates a new adapter instance
func AdapterConstructor(logger logging.Logger) (config.AdapterType, version.Adapter, error) {
	gh := github.NewClient(nil)
	return config.AdapterTypeGitHubRelease, newGitHubReleaseAdapter(logger, gh), nil
}
