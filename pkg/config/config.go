package config

import (
	"fmt"
)

// Config root object, just a list of targets
type Config struct {
	Targets []Target `yaml:"targets"`
}

// Target specifies a single target
type Target struct {
	Name    string        `yaml:"name"`
	Current AdapterConfig `yaml:"current"`
	Latest  AdapterConfig `yaml:"latest"`
}

// AdapterConfig config section
type AdapterConfig struct {
	Type              AdapterType       `yaml:"type"`
	K8sContainerImage K8sContainerImage `yaml:"k8sContainerImage,omitempty"`
	GitHubRelease     GitHubRelease     `yaml:"gitHubRelease"`
}

// K8sContainerImage config section
type K8sContainerImage struct {
	Kind          string `yaml:"kind"`
	Namespace     string `yaml:"namespace"`
	Name          string `yaml:"name"`
	ContainerName string `yaml:"containerName,omitempty"`
}

// String implements the fmt.Stringer interface
func (k *K8sContainerImage) String() string {
	return fmt.Sprintf("%v:%v/%v:%v", k.Kind, k.Namespace, k.Name, k.ContainerName)
}

// GitHubRelease config section
type GitHubRelease struct {
	Owner string `yaml:"owner"`
	Repo  string `yaml:"repo"`
}

// String implements the fmt.Stringer interface
func (k *GitHubRelease) String() string {
	return fmt.Sprintf("%v/%v", k.Owner, k.Repo)
}

// AdapterType defines the adapter to use to load the version from
type AdapterType string

// Current and Latest Adapter names
const (
	AdapterTypeK8sContainerImage AdapterType = "k8sContainerImage"
	AdapterTypeGitHubRelease     AdapterType = "gitHubRelease"
)
