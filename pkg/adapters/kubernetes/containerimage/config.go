package containerimage

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/version"
)

var (
	config envConfig
)

type envConfig struct {
	KubeContext string
	KubeConfig  string `envconfig:"KUBECONFIG"`
}

// AdapterConstructor creates a new adapter instance
func AdapterConstructor(logger logging.Logger) (version.Adapter, error) {
	envconfig.MustProcess("", &config)

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

// loadKubernetesClientConfig loads a REST Config as per the rules specified in GetConfig
// stolen from: https://github.com/kubernetes-sigs/controller-runtime/blob/2fe837fb5b0f4cfa9e566aa1027196a817692581/pkg/client/config/config.go
func loadKubernetesClientConfig() (*rest.Config, error) {
	if config.KubeConfig != "" {
		return loadConfigWithContext(config.KubeConfig)
	}

	if c, err := rest.InClusterConfig(); err == nil {
		return c, nil
	}

	if usr, err := user.Current(); err == nil {
		if c, err := loadConfigWithContext(filepath.Join(usr.HomeDir, ".kube", "config")); err == nil {
			return c, nil
		}
	}

	return nil, fmt.Errorf("could not locate a kubeconfig")
}

func loadConfigWithContext(kubeconfig string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{
			CurrentContext: config.KubeContext,
		},
	).ClientConfig()

}
