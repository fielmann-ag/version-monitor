package get

import (
	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// AdapterConstructor creates a new adapter instance
func AdapterConstructor(logger logging.Logger) (monitor.Adapter, error) {
	return newGetAdapter(logger), nil
}
