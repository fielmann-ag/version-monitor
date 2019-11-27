package html

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/fielmann-ag/version-monitor/pkg/internal/logging"
	"github.com/fielmann-ag/version-monitor/pkg/monitor"
)

// PageRenderer renders the fetched versions as a simple html page
type PageRenderer struct {
	monitor monitor.Monitor
	logger  logging.Logger
}

// NewPageRenderer returns a new PageRenderer
func NewPageRenderer(monitor monitor.Monitor, logger logging.Logger) *PageRenderer {
	return &PageRenderer{
		monitor: monitor,
		logger:  logger,
	}
}

func (r *PageRenderer) render(rw http.ResponseWriter) error {
	versions, date, err := r.monitor.Versions()
	if err != nil {
		return fmt.Errorf("failed to fetch versions from monitor: %v", err)
	}

	sort.Slice(versions, func(i, j int) bool { return versions[i].Name < versions[j].Name })

	params := &pageParams{
		Versions: versions,
		Date:     date.Format(time.RFC822),
	}
	if err := page.Execute(rw, params); err != nil {
		return fmt.Errorf("failed to render page template: %v", err)
	}

	return nil
}

// ServeHTTP implements the http.Handler interface
func (r *PageRenderer) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	if err := r.render(rw); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)

		r.logger.Errorf("failed to render versions template: %v", err)
		if _, errWrite := fmt.Fprintf(rw, "failed to render versions template: %v", err); errWrite != nil {
			r.logger.Errorf("error writing error message to response: %v", errWrite)
		}
	}
}
