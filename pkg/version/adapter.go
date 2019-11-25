package version

import (
	"fmt"
)

var adapters = map[string]Adapter{}

// AddAdapter adds an adapter with given unique name
func AddAdapter(name string, adapter Adapter) {
	if _, ok := adapters[name]; ok {
		panic(fmt.Sprintf("Adapter %s is already registered, please choose a different name", name))
	}
	adapters[name] = adapter
}
