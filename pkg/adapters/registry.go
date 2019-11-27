package adapters

import (
	"fmt"

	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// Registry of all loaded adapters
var Registry = map[monitor.AdapterType]monitor.Adapter{}

func register(name monitor.AdapterType, adapter monitor.Adapter) error {
	if _, ok := Registry[name]; ok {
		return fmt.Errorf("adapter %s is already registered", name)
	}

	Registry[name] = adapter
	return nil
}
