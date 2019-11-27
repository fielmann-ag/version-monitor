package adapters

import (
	"github.com/fielmann-ag/ops-version-monitor/pkg/config"
	"github.com/fielmann-ag/ops-version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/ops-version-monitor/pkg/version"
)

// AdapterConstructor constructs a new adapter
type AdapterConstructor func(logger logging.Logger) (config.AdapterType, version.Adapter, error)
