package containerimage

import (
	"errors"

	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// Kind types
const (
	KindDeployment  string = "Deployment"
	KindStatefulSet string = "StatefulSet"
	KindDaemonSet   string = "DaemonSet"

	// AdapterType constant
	AdapterType monitor.AdapterType = "k8sContainerImage"
)

var kinds = []string{
	KindDeployment,
	KindStatefulSet,
	KindDaemonSet,
}

// error codes for config violations
var (
	ErrNamespaceEmpty = errors.New("config has empty namespace field")
	ErrNameEmpty      = errors.New("config has empty name field")
)
