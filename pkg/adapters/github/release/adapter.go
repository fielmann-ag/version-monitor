package release

import (
	"github.com/fielmann-ag/ops-version-monitor/pkg/config"
	"github.com/fielmann-ag/ops-version-monitor/pkg/internal/logging"
)

type gitHubReleaseAdapter struct {
	logger       logging.Logger
	githubClient githubClient
}

func newGitHubReleaseAdapter(logger logging.Logger, githubClient githubClient) *gitHubReleaseAdapter {
	return &gitHubReleaseAdapter{
		logger:       logger,
		githubClient: githubClient,
	}
}

func (a *gitHubReleaseAdapter) Fetch(cfg config.AdapterConfig) (string, error) {
	return "", nil
}

func (a *gitHubReleaseAdapter) Validate(cfg config.AdapterConfig) error {
	if cfg.GitHubRelease.Repo == "" {
		return ErrReleaseEmpty
	}

	return nil
}
