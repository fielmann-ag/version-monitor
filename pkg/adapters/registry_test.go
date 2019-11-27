package adapters

import (
	"reflect"
	"testing"

	"github.com/fielmann-ag/version-monitor/pkg/config"
	testing2 "github.com/fielmann-ag/version-monitor/pkg/internal/testing"
)

func TestAddAdapter(t *testing.T) {
	var testAdapter1 = testing2.NewStaticAdapter("latest1")
	var testAdapter2 = testing2.NewStaticAdapter("latest2")

	type args struct {
		adapterType config.AdapterType
		adapter     Adapter
	}
	tests := []struct {
		name    string
		args    args
		before  map[config.AdapterType]Adapter
		after   map[config.AdapterType]Adapter
		wantErr bool
	}{
		{
			name: "add_simple",
			args: args{
				adapterType: "test2",
				adapter:     testAdapter2,
			},
			before: map[config.AdapterType]Adapter{
				"test1": testAdapter1,
			},
			after: map[config.AdapterType]Adapter{
				"test1": testAdapter1,
				"test2": testAdapter2,
			},
			wantErr: false,
		},
		{
			name: "add_simple",
			args: args{
				adapterType: "test1",
				adapter:     testAdapter1,
			},
			before: map[config.AdapterType]Adapter{
				"test1": testAdapter1,
			},
			after: map[config.AdapterType]Adapter{
				"test1": testAdapter1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Registry = tt.before
			err := register(tt.args.adapterType, tt.args.adapter)

			if (err != nil) != tt.wantErr {
				t.Errorf("AddAdapter() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.after, Registry) {
				t.Errorf("Expected to find adapter map %+v, but found %+v", tt.after, Registry)
			}
		})
	}
}
