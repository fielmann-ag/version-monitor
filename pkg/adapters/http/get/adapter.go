package get

import (
	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"net/http"
)

type getAdapter struct {
	logger logging.Logger
}

func newGetAdapter(logger logging.logger) *getAdapter {

	return &getAdapter{
		logger: logger,
	}
}

func (s getAdapter) Fetch(cfg monitor.AdapterConfig) (string, err) {
	resp.err := http.Get(cfg.httpGet.Url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL %v", err)
	}
	defer resp.Body.Close()

}
