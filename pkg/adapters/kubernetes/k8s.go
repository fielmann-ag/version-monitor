package kubernetes

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

// ContainerImageAdapter loads a version string from a k8s container image
type ContainerImageAdapter struct {
	logger    logging.Logger
	clientSet kubernetes.Interface
}

// NewContainerImageAdapter returns a new ContainerImageAdapter instance
func NewContainerImageAdapter(logger logging.Logger, clientSet kubernetes.Interface) *ContainerImageAdapter {
	return &ContainerImageAdapter{
		logger:    logger,
		clientSet: clientSet,
	}
}

func (a *ContainerImageAdapter) load(cfg config2.AdapterConfig) (*v1.PodTemplateSpec, error) {
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

// Validate the given config
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

// Fetch a version from given config
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
		if c.Name != cfg.K8sContainerImage.ContainerName {
			continue
		}

		a.logger.Debug(name + ": " + c.Image)
		return a.imageVersion(c), nil
	}

	return "", ErrContainerNotFound
}

func (a *ContainerImageAdapter) imageVersion(spec v1.Container) string {
	parts := strings.Split(spec.Image, ":")
	if len(parts) != 2 {
		return "latest"
	}
	return parts[1]
}
