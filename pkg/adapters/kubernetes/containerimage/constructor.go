package containerimage

import (
	"k8s.io/client-go/kubernetes"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// AdapterConstructor creates a new adapter instance
func AdapterConstructor(logger logging.Logger) (monitor.Adapter, error) {
	cfg, err := loadKubernetesClientConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	return newContainerImageAdapter(logger, client), nil
}
