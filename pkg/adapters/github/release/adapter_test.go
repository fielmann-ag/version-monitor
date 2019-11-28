package release

import (
	"errors"
	"testing"

	"github.com/google/go-github/v28/github"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

func Test_gitHubReleaseAdapter_Fetch(t *testing.T) {
	var testErr = errors.New("test error")

	type fields struct {
		reposClient reposClient
	}
	type args struct {
		cfg monitor.AdapterConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				reposClient: newTestRepoClient(nil, testErr),
			},
			args: args{
				cfg: monitor.AdapterConfig{
					GitHubRelease: monitor.GitHubRelease{
						Owner: "test",
						Repo:  "test",
					},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "simple",
			fields: fields{
				reposClient: newTestRepoClient(&github.RepositoryRelease{
					TagName: github.String("1.2.3"),
				}, nil),
			},
			args: args{
				cfg: monitor.AdapterConfig{
					GitHubRelease: monitor.GitHubRelease{
						Owner: "test",
						Repo:  "test",
					},
				},
			},
			want:    "1.2.3",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := newGitHubReleaseAdapter(logging.NewTestLogger(t), tt.fields.reposClient)
			got, err := a.Fetch(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Fetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gitHubReleaseAdapter_Validate(t *testing.T) {
	type args struct {
		cfg monitor.AdapterConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "no_repo_set",
			args: args{
				cfg: monitor.AdapterConfig{GitHubRelease: monitor.GitHubRelease{
					Owner: "test",
					Repo:  "",
				}},
			},
			wantErr: ErrRepoEmpty,
		},
		{
			name: "no_owner_set",
			args: args{
				cfg: monitor.AdapterConfig{GitHubRelease: monitor.GitHubRelease{
					Owner: "",
					Repo:  "test",
				}},
			},
			wantErr: ErrOwnerEmpty,
		},
		{
			name: "all_set",
			args: args{
				cfg: monitor.AdapterConfig{GitHubRelease: monitor.GitHubRelease{
					Owner: "test",
					Repo:  "test",
				}},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := newGitHubReleaseAdapter(nil, nil)
			if err := a.Validate(tt.args.cfg); err != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
