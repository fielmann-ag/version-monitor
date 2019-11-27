package adapters

import (
	"fmt"

	"github.com/fielmann-ag/ops-version-monitor/pkg/adapters/github/release"
	"github.com/fielmann-ag/ops-version-monitor/pkg/adapters/kubernetes/containerimage"
	"github.com/fielmann-ag/ops-version-monitor/pkg/internal/logging"
)

var adapters = []AdapterConstructor{
	containerimage.AdapterConstructor,
	release.AdapterConstructor,
}

func Register(logger logging.Logger) error {
	for _, constructor := range adapters {
		t, adapter, err := constructor(logger)
		if err != nil {
			return fmt.Errorf("failed to construct adapter %q: %v", t, err)
		}

		err = register(t, adapter)
		if err != nil {
			return fmt.Errorf("failed to register adapter %q: %v", t, err)
		}
	}

	return nil
}
