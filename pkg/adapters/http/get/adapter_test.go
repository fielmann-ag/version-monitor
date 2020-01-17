package get

import (
	"testing"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

func Test_getAdapter_Fetch(t *testing.T) {
	type fields struct {
		logger logging.Logger
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
			name:   "fine",
			fields: fields{},
			args: args{
				cfg: monitor.AdapterConfig{
					HttpGet: monitor.HttpGet{
						URL:      "https://ci.mgmt.ae.cloudhh.de/api/v1/info",
						JSONPath: "version",
					},
				},
			},
			want:    "5.7.2",
			wantErr: false,
		},
		{
			name:   "error",
			fields: fields{},
			args: args{
				cfg: monitor.AdapterConfig{
					HttpGet: monitor.HttpGet{
						URL:      "https://ci.mgmt.ae.cloudhh.de/api/v1/invalid",
						JSONPath: "version",
					},
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			s := &getAdapter{
				logger: tt.fields.logger,
			}

			got, err := s.Fetch(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAdapter.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getAdapter.Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}
