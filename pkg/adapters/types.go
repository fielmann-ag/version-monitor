package adapters

import (
	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// AdapterConstructor constructs a new adapter
type AdapterConstructor func(logger logging.Logger) (monitor.Adapter, error)
