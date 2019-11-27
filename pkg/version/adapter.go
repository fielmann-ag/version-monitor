package version

import (
	"fmt"

	"github.com/fielmann-ag/ops-version-monitor/pkg/config"
)

var adapters = map[config.AdapterType]Adapter{}

// AddAdapter adds an adapter with given unique name
func AddAdapter(name config.AdapterType, adapter Adapter) {
	if _, ok := adapters[name]; ok {
		panic(fmt.Sprintf("Adapter %s is already registered, please choose a different name", name))
	}
	adapters[name] = adapter
}
