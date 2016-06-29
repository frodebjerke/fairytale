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
	"github.com/frodebjerke/fairytale/handlers"
	"github.com/frodebjerke/fairytale/healthchecks"
	"github.com/frodebjerke/fairytale/storyteller"
	"github.com/udacity/ud615/app/health"
)

func main() {
	var (
		healthAddr = flag.String("health", "0.0.0.0:81", "Health service address.")
		httpAddr   = flag.String("http", "0.0.0.0:80", "Http service address.")
	)
	flag.Parse()

	log.Println("Starting server...")
	log.Printf("Health service listening on %s", *healthAddr)
	log.Printf("Http service listening on %s", *httpAddr)

	errChan := make(chan error, 10)

	healthchecks.NewServer(healthAddr, errChan)

	stories := storyteller.New()

	hmux := http.NewServeMux()
	hmux.HandleFunc("/verses", handlers.ReceiveDataHandler(stories))

	httpServer := manners.NewServer()
	httpServer.Addr = *httpAddr
	httpServer.Handler = hmux

	go func() {
		errChan <- httpServer.ListenAndServe()
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
