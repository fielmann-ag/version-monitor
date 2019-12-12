package get

import (
	"errors"

	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

const (
	AdapterType monitor.AdapterType = "httpGet"
)

var (
	ErrURLMissing = errors.New("httpGet.Url should be a URL")
)
