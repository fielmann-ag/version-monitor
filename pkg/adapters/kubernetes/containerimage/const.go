package containerimage

import (
	"errors"
)

// Kind types
const (
	KindDeployment  string = "Deployment"
	KindStatefulSet string = "StatefulSet"
	KindDaemonSet   string = "DaemonSet"
)

var kinds = []string{
	KindDeployment,
	KindStatefulSet,
	KindDaemonSet,
}

var (
	ErrNamespaceEmpty    = errors.New("config has empty namespace field")
	ErrNameEmpty         = errors.New("config has empty name field")
)
