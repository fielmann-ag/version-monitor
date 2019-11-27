package adapters

import (
	"reflect"
	"testing"

	testAdapter "github.com/fielmann-ag/version-monitor/pkg/internal/testing/adapter"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

func TestRegister(t *testing.T) {
	var testAdapter1 = testAdapter.NewStaticAdapter("latest1")
	var testAdapter2 = testAdapter.NewStaticAdapter("latest2")

	type args struct {
		adapterType monitor.AdapterType
		adapter     monitor.Adapter
	}
	tests := []struct {
		name    string
		args    args
		before  map[monitor.AdapterType]monitor.Adapter
		after   map[monitor.AdapterType]monitor.Adapter
		wantErr bool
	}{
		{
			name: "add_simple",
			args: args{
				adapterType: "test2",
				adapter:     testAdapter2,
			},
			before: map[monitor.AdapterType]monitor.Adapter{
				"test1": testAdapter1,
			},
			after: map[monitor.AdapterType]monitor.Adapter{
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
			before: map[monitor.AdapterType]monitor.Adapter{
				"test1": testAdapter1,
			},
			after: map[monitor.AdapterType]monitor.Adapter{
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
