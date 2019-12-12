package monitor

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
	GitHubRelease     GitHubRelease     `yaml:"gitHubRelease"`
	HttpGet           HttpGet           `yaml:"httpGet"`
	K8sContainerImage K8sContainerImage `yaml:"k8sContainerImage,omitempty"`
	ShellCommand      ShellCommand      `yaml:"shellCommand"`
}

// K8sContainerImage config section
type K8sContainerImage struct {
	Kind          string `yaml:"kind"`
	Namespace     string `yaml:"namespace"`
	Name          string `yaml:"name"`
	ContainerName string `yaml:"containerName,omitempty"`
}

// String implements the fmt.Stringer interface
func (k K8sContainerImage) String() string {
	return fmt.Sprintf("%v:%v/%v:%v", k.Kind, k.Namespace, k.Name, k.ContainerName)
}

// GitHubRelease config section
type GitHubRelease struct {
	Owner string `yaml:"owner"`
	Repo  string `yaml:"repo"`
}

// String implements the fmt.Stringer interface
func (g GitHubRelease) String() string {
	return fmt.Sprintf("%v/%v", g.Owner, g.Repo)
}

// ShellCommand config section
type ShellCommand struct {
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`
}

// String implements the fmt.Stringer interface
func (s ShellCommand) String() string {
	return fmt.Sprintf("%v", s.Command)
}

// HttpGet config section
type HttpGet struct {
	URL      string `yaml:"url"`
	JSONPath string `yaml:"jsonPath"`
}

// String implements the fmt.Stringer interface
func (h HttpGet) String() string {
	return h.URL
}
