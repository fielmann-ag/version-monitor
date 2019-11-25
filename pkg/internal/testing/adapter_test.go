package testing

import (
	"testing"
)

func TestStaticAdapter_FetchVersion(t *testing.T) {
	type fields struct {
		Value string
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
				Value: "test-value",
			},
			want:    "test-value",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &StaticAdapter{
				Value: tt.fields.Value,
			}
			got, err := a.FetchVersion()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FetchVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}
