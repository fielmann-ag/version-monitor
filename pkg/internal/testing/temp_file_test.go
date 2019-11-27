package testing

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestTempFile(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "simple",
			args: args{
				content: `smalll content
within file`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := TempFile(t, tt.args.content)

			if _, err := os.Stat(got); err != nil {
				t.Fatalf("expected to get valid filepath and existing file at %v, got er: %v", got, err)
			}

			if content, err := ioutil.ReadFile(got); err != nil {
				t.Errorf("failed to read file %v: %v", got, err)
			}  else if string(content) != tt.args.content {
				t.Errorf("exepcted to find content %q, but found %q", tt.args.content, string(content))
			}

			got1()

			if _, err := os.Stat(got); err != nil && !os.IsNotExist(err) {
				t.Fatalf("expected to get valid filepath and existing file at %v, got er: %v", got, err)
			}
		})
	}
}
