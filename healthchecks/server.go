package healthchecks

import (
	"net/http"

	"github.com/braintree/manners"
	"github.com/udacity/ud615/app/health"
)

// NewServer creates a new health and readiness server
func NewServer(addr *string, errChan chan error) {
	hmux := http.NewServeMux()
	hmux.HandleFunc("/healthz", health.HealthzHandler)
	hmux.HandleFunc("/readiness", health.ReadinessHandler)
	hmux.HandleFunc("/healthz/status", health.HealthzStatusHandler)
	hmux.HandleFunc("/readiness/status", health.ReadinessStatusHandler)

	healthServer := manners.NewServer()
	healthServer.Addr = *addr

	go func() {
		errChan <- healthServer.ListenAndServe()
	}()
}
