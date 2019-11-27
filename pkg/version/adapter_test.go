package version

import (
	"reflect"
	"testing"

	"github.com/fielmann-ag/ops-version-monitor/pkg/config"
	testing2 "github.com/fielmann-ag/ops-version-monitor/pkg/internal/testing"
)

func TestAddAdapter(t *testing.T) {
	var testAdapter1 = testing2.NewStaticAdapter("test1", "test1", "latest1")
	var testAdapter2 = testing2.NewStaticAdapter("test1", "test2", "latest2")

	type args struct {
		adapterType config.AdapterType
		adapter     Adapter
	}
	tests := []struct {
		name      string
		args      args
		before    map[config.AdapterType]Adapter
		after     map[config.AdapterType]Adapter
		wantPanic bool
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
			wantPanic: false,
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
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("AddAdapter() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()

			adapters = tt.before
			AddAdapter(tt.args.adapterType, tt.args.adapter)

			if !reflect.DeepEqual(tt.after, adapters) {
				t.Errorf("Expected to find adapter map %+v, but found %+v", tt.after, adapters)
			}
		})
	}
}
