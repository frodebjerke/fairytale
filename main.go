package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/udacity/ud615/app/handlers"
	"github.com/udacity/ud615/app/health"
)

func main() {
	var (
		healthAddr = flag.String("health", "0.0.0.0:81", "Health service address.")
	)
	flag.Parse()

	log.Println("Starting server...")
	log.Printf("Health service listening on %s", *healthAddr)

	errChan := make(chan error, 10)

	hmux := http.NewServeMux()
	hmux.HandleFunc("/healthz", health.HealthzHandler)
	hmux.HandleFunc("/readiness", health.ReadinessHandler)
	hmux.HandleFunc("/healthz/status", health.HealthzStatusHandler)
	hmux.HandleFunc("/readiness/status", health.ReadinessStatusHandler)

	healthServer := manners.NewServer()
	healthServer.Addr = *healthAddr
	healthServer.Handler = handlers.LoggingHandler(hmux)

	go func() {
		errChan <- healthServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			health.SetReadinessStatus(http.StatusServiceUnavailable)
			os.Exit(0)
		}
	}
}
