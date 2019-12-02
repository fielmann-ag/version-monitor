package version

import (
	"regexp"
)

var cleanRegExp = regexp.MustCompile(`[0-9]+(?:\.[0-9]+)+`)

// Clean removes all string values from given version string
func Clean(version string) string {
	return cleanRegExp.FindString(version)
}
