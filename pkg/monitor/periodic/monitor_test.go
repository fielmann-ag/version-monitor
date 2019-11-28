package periodic

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	internalTesting "github.com/fielmann-ag/version-monitor/pkg/internal/testing"
	"github.com/fielmann-ag/version-monitor/pkg/internal/testing/adapter"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

var errTest = errors.New("test")

func TestPeriodicMonitor_Start(t *testing.T) {
	type fields struct {
		config   *monitor.Config
		adapters monitor.AdapterRegistry
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "empty_config",
			fields: fields{
				config:   &monitor.Config{Targets: []monitor.Target{}},
				adapters: nil,
			},
			wantErr: errors.New("no targets defined"),
		},
		{
			name: "valid_config",
			fields: fields{
				config: &monitor.Config{
					Targets: []monitor.Target{
						{
							Name: "static",
							Current: monitor.AdapterConfig{
								Type: adapter.AdapterTypeStatic,
							},
							Latest: monitor.AdapterConfig{
								Type: adapter.AdapterTypeStatic,
							},
						},
					},
				},
				adapters: map[monitor.AdapterType]monitor.Adapter{
					adapter.AdapterTypeStatic: adapter.NewStaticAdapter("1.2.3"),
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMonitor(logging.NewTestLogger(t), tt.fields.config, tt.fields.adapters)
			err := m.Start()

			if (err != nil) != (tt.wantErr != nil) || (err != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPeriodicMonitor_Versions(t *testing.T) {
	type fields struct {
		config   *monitor.Config
		adapters monitor.AdapterRegistry

		cachedVersions   map[string]monitor.Version
		lastError        error
		latestResultFrom time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    []monitor.Version
		want1   time.Time
		wantErr bool
	}{
		{
			name: "simple",
			fields: fields{
				config:   nil,
				adapters: nil,
				cachedVersions: map[string]monitor.Version{
					"test": {
						Name:    "test",
						Current: "1.2.3",
						Latest:  "1.2.5",
					},
				},
				lastError:        nil,
				latestResultFrom: internalTesting.Time,
			},
			want: []monitor.Version{
				{
					Name:    "test",
					Current: "1.2.3",
					Latest:  "1.2.5",
				},
			},
			want1:   internalTesting.Time,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				config:           nil,
				adapters:         nil,
				cachedVersions:   nil,
				lastError:        errTest,
				latestResultFrom: time.Time{},
			},
			want:    nil,
			want1:   time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Monitor{
				cachedVersions:   tt.fields.cachedVersions,
				lastError:        tt.fields.lastError,
				latestResultFrom: tt.fields.latestResultFrom,
			}

			got, got1, err := m.Versions()
			if (err != nil) != tt.wantErr {
				t.Errorf("Versions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Versions() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Versions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPeriodicMonitor_error(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "error",
			args: args{
				err: errTest,
			},
		},
		{
			name: "nil",
			args: args{
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMonitor(logging.NewTestLogger(t), nil, nil)
			m.error(tt.args.err)
			if (m.lastError != nil) != (tt.args.err != nil) || (m.lastError != nil && m.lastError.Error() != tt.args.err.Error()) {
				t.Errorf("error() error = %v, wantErr %v", m.lastError, tt.args.err)
			}
		})
	}
}

func TestPeriodicMonitor_storeVersion(t *testing.T) {
	type args struct {
		targetName string
		version    monitor.Version
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "simple",
			args: args{
				targetName: "test-1",
				version: monitor.Version{
					Name:    "test-1",
					Current: "1.2.3",
					Latest:  "2.3.4",
				},
			},
		},
		{
			name: "empty",
			args: args{
				targetName: "",
				version:    monitor.Version{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMonitor(logging.NewTestLogger(t), nil, nil)
			if len(m.cachedVersions) != 0 {
				t.Fatalf("Expected to have 0 cachedVersions, found %v", len(m.cachedVersions))
			}

			m.storeVersion(tt.args.targetName, tt.args.version)

			if len(m.cachedVersions) != 1 {
				t.Fatalf("Expected to have 1 cachedVersions, found %v", len(m.cachedVersions))
			}

		})
	}
}

func TestPeriodicMonitor_fetch(t *testing.T) {
	type args struct {
		target monitor.Target
	}
	tests := []struct {
		name     string
		adapters map[monitor.AdapterType]monitor.Adapter
		args     args
	}{
		{
			name: "static",
			adapters: map[monitor.AdapterType]monitor.Adapter{
				adapter.AdapterTypeStatic: adapter.NewStaticAdapter("1.2.3"),
			},
			args: args{
				target: monitor.Target{
					Name: "static",
					Current: monitor.AdapterConfig{
						Type: adapter.AdapterTypeStatic,
					},
					Latest: monitor.AdapterConfig{
						Type: adapter.AdapterTypeStatic,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMonitor(logging.NewTestLogger(t), nil, tt.adapters)
			if len(m.cachedVersions) != 0 {
				t.Errorf("expected to find 0 cachedVersion, found %v", len(m.cachedVersions))
			}
			m.fetch(tt.args.target)
			if len(m.cachedVersions) != 1 {
				t.Errorf("expected to find 1 cachedVersion, found %v", len(m.cachedVersions))
			}
		})
	}
}

func TestPeriodicMonitor_validateConfig(t *testing.T) {
	type fields struct {
		config   *monitor.Config
		adapters monitor.AdapterRegistry
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "nil",
			fields: fields{
				config:   nil,
				adapters: nil,
			},
			wantErr: errors.New("config is not set"),
		},
		{
			name: "empty_targets",
			fields: fields{
				config: &monitor.Config{
					Targets: []monitor.Target{},
				},
				adapters: nil,
			},
			wantErr: errors.New("no targets defined"),
		},
		{
			name: "target_current_not_exists",
			fields: fields{
				config: &monitor.Config{
					Targets: []monitor.Target{
						monitor.Target{
							Name: "static",
							Current: monitor.AdapterConfig{
								Type: "not_exists",
							},
							Latest: monitor.AdapterConfig{
								Type: adapter.AdapterTypeStatic,
							},
						},
					},
				},
				adapters: map[monitor.AdapterType]monitor.Adapter{
					adapter.AdapterTypeStatic: adapter.NewStaticAdapter("1.2.3.4"),
				},
			},
			wantErr: errors.New("target.current.type \"not_exists\" of target static not found"),
		},
		{
			name: "target_latest_not_exists",
			fields: fields{
				config: &monitor.Config{
					Targets: []monitor.Target{
						monitor.Target{
							Name: "static",
							Current: monitor.AdapterConfig{
								Type: adapter.AdapterTypeStatic,
							},
							Latest: monitor.AdapterConfig{
								Type: "not_exists",
							},
						},
					},
				},
				adapters: map[monitor.AdapterType]monitor.Adapter{
					adapter.AdapterTypeStatic: adapter.NewStaticAdapter("1.2.3.4"),
				},
			},
			wantErr: errors.New("target.latest.type \"not_exists\" of target static not found"),
		},
		{
			name: "single_target",
			fields: fields{
				config: &monitor.Config{
					Targets: []monitor.Target{
						{
							Name: "static",
							Current: monitor.AdapterConfig{
								Type: adapter.AdapterTypeStatic,
							},
							Latest: monitor.AdapterConfig{
								Type: adapter.AdapterTypeStatic,
							},
						},
					},
				},
				adapters: map[monitor.AdapterType]monitor.Adapter{
					adapter.AdapterTypeStatic: adapter.NewStaticAdapter("1.2.3.4"),
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMonitor(logging.NewTestLogger(t), tt.fields.config, tt.fields.adapters)

			err := m.validateConfig()
			if (err != nil) != (tt.wantErr != nil) || (err != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
