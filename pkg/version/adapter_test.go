package version

import (
	"reflect"
	"testing"

	testing2 "github.com/fielmann-ag/ops-version-monitor/pkg/internal/testing"
)

func TestAddAdapter(t *testing.T) {
	var testAdapter1 = testing2.NewStaticAdapter("test1")
	var testAdapter2 = testing2.NewStaticAdapter("test2")

	type args struct {
		name    string
		adapter Adapter
	}
	tests := []struct {
		name      string
		args      args
		before    map[string]Adapter
		after     map[string]Adapter
		wantPanic bool
	}{
		{
			name: "add_simple",
			args: args{
				name:    "test2",
				adapter: testAdapter2,
			},
			before: map[string]Adapter{
				"test1": testAdapter1,
			},
			after: map[string]Adapter{
				"test1": testAdapter1,
				"test2": testAdapter2,
			},
			wantPanic: false,
		},
		{
			name: "add_simple",
			args: args{
				name:    "test1",
				adapter: testAdapter1,
			},
			before: map[string]Adapter{
				"test1": testAdapter1,
			},
			after: map[string]Adapter{
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
			AddAdapter(tt.args.name, tt.args.adapter)

			if !reflect.DeepEqual(tt.after, adapters) {
				t.Errorf("Expected to find adapter map %+v, but found %+v", tt.after, adapters)
			}
		})
	}
}
