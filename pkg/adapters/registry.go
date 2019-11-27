package adapters

import (
	"fmt"

	"github.com/fielmann-ag/version-monitor/pkg/config"
	"github.com/fielmann-ag/version-monitor/pkg/version"
)

// Registry of all loaded adapters
var Registry = map[config.AdapterType]version.Adapter{}

func register(name config.AdapterType, adapter version.Adapter) error {
	if _, ok := Registry[name]; ok {
		return fmt.Errorf("adapter %s is already registered", name)
	}

	Registry[name] = adapter
	return nil
}
