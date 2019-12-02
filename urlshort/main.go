package main

import (
	"context"
	"flag"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/skycoop/gophercises/urlshort/handlers"
	"github.com/skycoop/gophercises/urlshort/urlregistry"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func serve(srv *http.Server) {
	err := srv.ListenAndServe()
	if err != nil {
		log.WithField("error", err).Error("Server stopped serving with error")
	}
}

func main() {
	address := flag.String("address", ":8080", "The address for the server to listen to")
	debug := flag.Bool("debug", false, "Whether to output debug logging")
	gracefulTimeout := flag.Duration("graceful-timeout", 15*time.Second,
		"The max time the server will wait for a graceful shutdown")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	registry := urlregistry.NewMemoryUrlRegistry()

	router := mux.NewRouter()
	router.Handle("/register", handlers.NewRegisterHandler(registry)).
		Methods("POST").
		Headers("Content-Type", "application/json").
		Name("register")
	router.Handle("/{guid:[0-9a-v]{20}}", handlers.NewLookupHandler(registry)).
		Methods("GET").
		Name("lookup")

	srv := &http.Server{
		Addr:         *address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	log.Info("Starting server")
	go serve(srv)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	log.Info("Interrupt received, attempting graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), *gracefulTimeout)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.WithField("error", err).Fatal("Error while attempting graceful shutdown")
	} else {
		log.Info("Server shut down gracefully")
	}
}
