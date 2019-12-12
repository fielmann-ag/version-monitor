package get

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

var _ monitor.Adapter = &getAdapter{}

type getAdapter struct {
	logger logging.Logger
}

func newGetAdapter(logger logging.Logger) *getAdapter {
	return &getAdapter{
		logger: logger,
	}
}

func (s *getAdapter) Fetch(cfg monitor.AdapterConfig) (string, error) {
	res, err := http.Get(cfg.HttpGet.URL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL %v: %v", cfg.HttpGet.URL, err)

	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if err := res.Body.Close(); err != nil {
		return "", fmt.Errorf("failed to close response body: %v", err)
	}

	if !gjson.ValidBytes(content) {
		return "", fmt.Errorf("invalid json body from url %v: %v", cfg.HttpGet.URL, string(content))
	}

	value := gjson.Get(string(content), cfg.HttpGet.JSONPath)

	return value.String(), nil

}

func (s *getAdapter) Validate(cfg monitor.AdapterConfig) error {
	if cfg.HttpGet.URL == "" {
		return ErrURLMissing
	}
	return nil
}
