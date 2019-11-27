package testing

import (
	"testing"

	"github.com/fielmann-ag/ops-version-monitor/pkg/config"
)

func TestStaticAdapter_FetchVersion(t *testing.T) {
	type fields struct {
		Version string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "same_result",
			fields: fields{
				Version: "test-value",
			},
			want:    "test-value",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &StaticAdapter{
				Version: tt.fields.Version,
			}
			got, err := a.Fetch(config.AdapterConfig{})
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
