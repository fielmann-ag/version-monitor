package get

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

//var _ monitor.Adapter = &shellCommandAdapter{}

var _ monitor.Adapter = &getAdapter{}

type getAdapter struct {
	logger logging.Logger
}

func newGetAdapter(logger logging.Logger) *getAdapter {
	return &getAdapter{
		logger: logger,
	}
}

func (s getAdapter) Fetch(cfg monitor.AdapterConfig) (string, error) {

	res, err := http.Get(cfg.HttpGet.URL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL %v", err)

	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL %v", err)
	}
	return content.string, nil

}

func (s getAdapter) Validate(cfg monitor.AdapterConfig) error {
	if cfg.HttpGet.URL == "" {
		return ErrURLMissing
	}
	return nil
}
