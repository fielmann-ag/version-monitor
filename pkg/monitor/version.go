package monitor

import (
	"regexp"
)

var cleanRegExp = regexp.MustCompile("[^0-9\\\\.]+")

func cleanVersion(version string) string {
	return cleanRegExp.ReplaceAllString(version, "")
}
