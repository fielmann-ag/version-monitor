package release

import (
	"context"
	"fmt"

	"github.com/fielmann-ag/ops-version-monitor/pkg/config"
	"github.com/fielmann-ag/ops-version-monitor/pkg/internal/logging"
)

type gitHubReleaseAdapter struct {
	logger      logging.Logger
	reposClient reposClient
}

func newGitHubReleaseAdapter(logger logging.Logger, reposClient reposClient) *gitHubReleaseAdapter {
	return &gitHubReleaseAdapter{
		logger:      logger,
		reposClient: reposClient,
	}
}

func (a *gitHubReleaseAdapter) Fetch(cfg config.AdapterConfig) (string, error) {
	a.logger.Debugf("Fetching latest release from %v", cfg.GitHubRelease)
	release, _, err := a.reposClient.GetLatestRelease(context.TODO(), cfg.GitHubRelease.Owner, cfg.GitHubRelease.Repo)
	if err != nil {
		return "", fmt.Errorf("failed to load GH repo %v: %v", cfg.GitHubRelease, err)
	}

	return *release.TagName, nil
}

func (a *gitHubReleaseAdapter) Validate(cfg config.AdapterConfig) error {
	if cfg.GitHubRelease.Owner == "" {
		return ErrOwnerEmpty
	}
	if cfg.GitHubRelease.Repo == "" {
		return ErrRepoEmpty
	}

	return nil
}
