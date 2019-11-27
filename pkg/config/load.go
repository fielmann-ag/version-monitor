package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// Load the config from a given filename (including path)
func Load(filename string) (*monitor.Config, error) {
	if err := exists(filename); err != nil {
		return nil, fmt.Errorf("failed to load file %v: %v", filename, err)
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load file %v: %v", filename, err)
	}

	cfg := &monitor.Config{}
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file %v: %v", filename, err)
	}

	return cfg, nil
}

func exists(filename string) error {
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return fmt.Errorf("file %s is a directoy", filename)
	}

	if info.Size() == 0 {
		return fmt.Errorf("file %s is empty", filename)
	}

	return nil
}
