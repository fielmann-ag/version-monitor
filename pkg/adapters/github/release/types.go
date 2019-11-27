package release

import (
	"context"

	"github.com/google/go-github/v28/github"
)

type reposClient interface {
	GetLatestRelease(ctx context.Context, owner, repo string) (*github.RepositoryRelease, *github.Response, error)
}
