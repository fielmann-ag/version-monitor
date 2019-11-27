package release

import (
	"github.com/google/go-github/v28/github"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// AdapterConstructor creates a new adapter instance
func AdapterConstructor(logger logging.Logger) (monitor.Adapter, error) {
	gh := github.NewClient(nil)
	return newGitHubReleaseAdapter(logger, gh.Repositories), nil
}
