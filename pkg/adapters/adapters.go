package adapters

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/fielmann-ag/version-monitor/pkg/adapters/github/release"
	"github.com/fielmann-ag/version-monitor/pkg/adapters/kubernetes/containerimage"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

var adapters = map[monitor.AdapterType]AdapterConstructor{
	monitor.AdapterTypeK8sContainerImage: containerimage.AdapterConstructor,
	monitor.AdapterTypeGitHubRelease:     release.AdapterConstructor,
}

// Register all adapters using their constructors
func Register(logger *logrus.Logger) error {
	for t, constructor := range adapters {
		adapter, err := constructor(logger.WithField("adapter", t))
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
