package release

import (
	"context"
	"fmt"

	"github.com/google/go-github/v28/github"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

var _ monitor.Adapter = &gitHubReleaseAdapter{}

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

func (a *gitHubReleaseAdapter) Fetch(cfg monitor.AdapterConfig) (string, error) {
	a.logger.Debugf("Fetching latest release from %v", cfg.GitHubRelease)
	release, _, err := a.reposClient.GetLatestRelease(context.TODO(), cfg.GitHubRelease.Owner, cfg.GitHubRelease.Repo)
	if err != nil {
		// if we're able to cast the error AND it is a not found there are multiple reasons for that. The docs at
		// https://developer.github.com/v3/repos/releases/#get-the-latest-release state that the response contains the
		// `most recent non-prerelease, non-draft release` which could also mean that the releases are marked as draft
		// or not marked as stable. 🤷🏼‍
		if ghErr, ok := err.(*github.ErrorResponse); ok && ghErr.Message == "Not Found" {
			return "", nil
		}

		return "", fmt.Errorf("failed to load GH repo %v: %v", cfg.GitHubRelease, err)
	}

	return *release.TagName, nil
}

func (a *gitHubReleaseAdapter) Validate(cfg monitor.AdapterConfig) error {
	if cfg.GitHubRelease.Owner == "" {
		return ErrOwnerEmpty
	}
	if cfg.GitHubRelease.Repo == "" {
		return ErrRepoEmpty
	}

	return nil
}
