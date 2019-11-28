package release

import (
	"context"

	"github.com/google/go-github/v28/github"
)

type testRepoClient struct {
	err     error
	release *github.RepositoryRelease
}

func newTestRepoClient(release *github.RepositoryRelease, err error) *testRepoClient {
	return &testRepoClient{
		err:     err,
		release: release,
	}
}

func (c *testRepoClient) GetLatestRelease(ctx context.Context, owner, repo string) (*github.RepositoryRelease, *github.Response, error) {
	return c.release, nil, c.err
}
