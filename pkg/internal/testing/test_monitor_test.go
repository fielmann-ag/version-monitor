package testing

import (
	"reflect"
	"testing"
	"time"

	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

func TestMonitor_Versions(t *testing.T) {
	type fields struct {
		versions []monitor.Version
		t        time.Time
		err      error
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
				versions: []monitor.Version{
					{Name: "x-test", Current: "1.2.3", Latest: "1.4.5"},
				},
				t:   Time,
				err: nil,
			},
			want: []monitor.Version{
				{Name: "x-test", Current: "1.2.3", Latest: "1.4.5"},
			},
			want1:   Time,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMonitor(tt.fields.versions, tt.fields.t, tt.fields.err)
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
