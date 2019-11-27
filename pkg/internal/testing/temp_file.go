package testing

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// TempFile returns a new temp file with given context and returns the filepath and a delete func
func TempFile(t *testing.T, content string) (string, func()) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "test-")
	if err != nil {
		t.Fatalf("Cannot create temporary file: %v", err)
	}

	if _, err := fmt.Fprint(tmpFile, content); err != nil {
		t.Fatalf("failed to write tmp content to file: %v", err)
	}

	if err := tmpFile.Close(); err != nil {
		t.Fatalf("failed to close tmp file: %v", err)
	}

	return tmpFile.Name(), func() {
		if err := os.Remove(tmpFile.Name()); err != nil {
			t.Fatalf("failed to remove tmp file: %v", err)
		}
	}
}
