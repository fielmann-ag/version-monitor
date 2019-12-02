package version

import (
	"testing"
)

func TestClean(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "noop",
			args: args{
				version: "1.2.3",
			},
			want: "1.2.3",
		},
		{
			name: "trim_left",
			args: args{
				version: "test-1.2.3",
			},
			want: "1.2.3",
		},
		{
			name: "trim_right",
			args: args{
				version: "1.2.3-test",
			},
			want: "1.2.3",
		},
		{
			name: "trim_both",
			args: args{
				version: "test-1.2.3-test",
			},
			want: "1.2.3",
		},
		{
			name: "non_semver",
			args: args{
				version: "1.2.3.4.5",
			},
			want: "1.2.3.4.5",
		},
		{
			name: "dot_postfix",
			args: args{
				version: "1.2.3.4.5.",
			},
			want: "1.2.3.4.5",
		},
		{
			name: "non_semver_multi_digit",
			args: args{
				version: "11.22.33.44.55",
			},
			want: "11.22.33.44.55",
		},
		{
			name: "dot_postfix_multi_digit",
			args: args{
				version: "11.22.33.44.55.",
			},
			want: "11.22.33.44.55",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clean(tt.args.version); got != tt.want {
				t.Errorf("Clean() = %v, want %v", got, tt.want)
			}
		})
	}
}
