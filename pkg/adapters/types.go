package adapters

import (
	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/version"
)

// AdapterConstructor constructs a new adapter
type AdapterConstructor func(logger logging.Logger) (version.Adapter, error)