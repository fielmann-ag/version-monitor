package command

import (
	"errors"

	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// Adapter Error
var (
	ErrCommandEmpty = errors.New("shellCommand.command must not be empty")
)
// AdapterType constant
const AdapterType monitor.AdapterType = "shellCommand"
