package get

import (
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
	resp, err := http.Get(cfg
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Reading URL failed%v", err)
	}
	*/
	return "", nil

}

func (s getAdapter) Validate(cfg monitor.AdapterConfig) error {
	/* if cfg.HtttpGet == "" {

	}
	*/
	//	if cfg. == "" {
	//		return ErrCommandEmpty
	//	}

	return nil
}
