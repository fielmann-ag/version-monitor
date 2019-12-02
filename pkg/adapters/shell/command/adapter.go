package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

var _ monitor.Adapter = &shellCommandAdapter{}

type shellCommandAdapter struct {
	logger logging.Logger
}

func newShellCommandAdapter(logger logging.Logger) *shellCommandAdapter {
	return &shellCommandAdapter{
		logger: logger,
	}
}

func (s shellCommandAdapter) Fetch(cfg monitor.AdapterConfig) (string, error) {
	cmd := exec.Command(cfg.ShellCommand.Command, cfg.ShellCommand.Args...)
	cmd.Env = os.Environ()

	var err error
	if cmd.Dir, err = os.Getwd(); err != nil {
		return "", fmt.Errorf("failed to get CWD: %v", err)
	}

	resultBytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run %v: %v", cfg.ShellCommand.Command, err)
	}

	return string(resultBytes), nil
}

func (s shellCommandAdapter) Validate(cfg monitor.AdapterConfig) error {
	if cfg.ShellCommand.Command == "" {
		return ErrCommandEmpty
	}

	return nil
}
