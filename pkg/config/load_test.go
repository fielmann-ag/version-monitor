package config

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/fielmann-ag/version-monitor/pkg/adapters/github/release"
	"github.com/fielmann-ag/version-monitor/pkg/adapters/kubernetes/containerimage"
	testing2 "github.com/fielmann-ag/version-monitor/pkg/internal/testing"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name        string
		fileExists  bool
		fileContent string
		want        *monitor.Config
		wantErr     bool
	}{
		{
			name:        "file_not_exists",
			fileExists:  false,
			fileContent: "",
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "empty_file",
			fileExists:  true,
			fileContent: "",
			want:        nil,
			wantErr:     true,
		},
		{
			name:        "no_targets",
			fileExists:  true,
			fileContent: `targets: []`,
			want:        &monitor.Config{Targets: []monitor.Target{}},
			wantErr:     false,
		},
		{
			name:       "single_target",
			fileExists: true,
			fileContent: `
targets:
  - name: static
    current:
      type: static
    latest:
      type: static
`,
			want: &monitor.Config{
				Targets: []monitor.Target{
					{
						Name: "static",
						Current: monitor.AdapterConfig{
							Type: "static",
						},
						Latest: monitor.AdapterConfig{
							Type: "static",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:       "multiple_target",
			fileExists: true,
			fileContent: `
targets:
  - name: full
    current:
      type: k8sContainerImage
      k8sContainerImage:
        kind: Deployment
        namespace: testing
        name: test
        containerName: test-container
    latest:
      type: gitHubRelease
      gitHubRelease:
        owner: test-owner
        repo: test-repo
`,
			want: &monitor.Config{
				Targets: []monitor.Target{
					{
						Name: "full",
						Current: monitor.AdapterConfig{
							Type: containerimage.AdapterType,
							K8sContainerImage: monitor.K8sContainerImage{
								Kind:          "Deployment",
								Namespace:     "testing",
								Name:          "test",
								ContainerName: "test-container",
							},
						},
						Latest: monitor.AdapterConfig{
							Type: release.AdapterType,
							GitHubRelease: monitor.GitHubRelease{
								Owner: "test-owner",
								Repo:  "test-repo",
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filename string
			var deleteFunc func()

			if tt.fileExists {
				filename, deleteFunc = testing2.TempFile(t, tt.fileContent)
				defer deleteFunc()
			} else {
				filename = fmt.Sprintf("/tmp/%v", rand.Uint64())
			}

			got, err := Load(filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exists(t *testing.T) {
	tests := []struct {
		name    string
		exists  bool
		wantErr bool
	}{
		{
			name:    "exists",
			exists:  true,
			wantErr: false,
		},
		{
			name:    "not_exists",
			exists:  false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filename string
			var deleteFunc func()

			if tt.exists {
				filename, deleteFunc = testing2.TempFile(t, "something")
				defer deleteFunc()
			} else {
				filename = fmt.Sprintf("/tmp/%v", rand.Uint64())
			}

			if err := exists(filename); (err != nil) != tt.wantErr {
				t.Errorf("exists() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
