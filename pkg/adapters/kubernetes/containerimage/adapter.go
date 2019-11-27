package containerimage

import (
	"fmt"
	"strings"

	"github.com/apinnecke/go-stringslice"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	config2 "github.com/fielmann-ag/ops-version-monitor/pkg/config"
	"github.com/fielmann-ag/ops-version-monitor/pkg/internal/logging"
)

type ContainerImageAdapter struct {
	logger    logging.Logger
	clientSet kubernetes.Interface
}

func newContainerImageAdapter(logger logging.Logger, clientSet kubernetes.Interface) *ContainerImageAdapter {
	return &ContainerImageAdapter{
		logger:    logger,
		clientSet: clientSet,
	}
}

func (a *ContainerImageAdapter) load(cfg config2.AdapterConfig) (*v1.PodTemplateSpec, error) {
	a.logger.Debugf("Loading kubernetes %v", cfg.K8sContainerImage)

	if cfg.K8sContainerImage.Kind == KindDeployment {
		dep, err := a.clientSet.AppsV1().Deployments(cfg.K8sContainerImage.Namespace).Get(cfg.K8sContainerImage.Name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to load Deployment: %v", err)
		}

		return &dep.Spec.Template, nil
	}

	if cfg.K8sContainerImage.Kind == KindDaemonSet {
		dep, err := a.clientSet.AppsV1().DaemonSets(cfg.K8sContainerImage.Namespace).Get(cfg.K8sContainerImage.Name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to load DaemonSet: %v", err)
		}

		return &dep.Spec.Template, nil
	}

	if cfg.K8sContainerImage.Kind == KindStatefulSet {
		dep, err := a.clientSet.AppsV1().StatefulSets(cfg.K8sContainerImage.Namespace).Get(cfg.K8sContainerImage.Name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to load StatefulSet: %v", err)
		}

		return &dep.Spec.Template, nil
	}

	panic("Dead code reached. Receiving a kind that is unknown must be prevented by Validate() method.")
}

func (a *ContainerImageAdapter) Validate(cfg config2.AdapterConfig) error {
	if !stringslice.Contains(kinds, cfg.K8sContainerImage.Kind) {
		return fmt.Errorf("invalid k8sContainerImage.kind value %s (valid are %v)", cfg.K8sContainerImage.Kind, kinds)
	}
	if cfg.K8sContainerImage.Namespace == "" {
		return ErrNamespaceEmpty
	}
	if cfg.K8sContainerImage.Name == "" {
		return ErrNameEmpty
	}

	return nil
}

func (a *ContainerImageAdapter) Fetch(cfg config2.AdapterConfig) (string, error) {
	podTemplate, err := a.load(cfg)
	if err != nil {
		return "", err
	}

	name := cfg.K8sContainerImage.Namespace + cfg.K8sContainerImage.Name

	if cfg.K8sContainerImage.ContainerName == "" && len(podTemplate.Spec.Containers) > 1 {
		return "", fmt.Errorf("failed to load version from k8s %s: PodTemplate has more then 1 container but no ContainerName is specified", name)
	}

	for _, c := range podTemplate.Spec.Containers {
		if cfg.K8sContainerImage.ContainerName != "" && cfg.K8sContainerImage.ContainerName != c.Name {
			continue
		}

		a.logger.Debug(name + ": " + c.Image)
		return a.imageVersion(c), nil
	}

	return "", fmt.Errorf("podTemplate of %s does not have the desired container", cfg.K8sContainerImage)
}

func (a *ContainerImageAdapter) imageVersion(spec v1.Container) string {
	parts := strings.Split(spec.Image, ":")
	if len(parts) != 2 {
		return "latest"
	}
	return parts[1]
}
