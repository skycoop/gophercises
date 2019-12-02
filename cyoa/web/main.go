package main

import (
	"context"
	"flag"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/skycoop/gophercises/cyoa"
	"github.com/skycoop/gophercises/cyoa/web/handlers"
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
	storyFile := flag.String("story-file", "cyoa/data/gopher.json", "The JSON file to load the story from")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	story, err := cyoa.LoadStory(*storyFile)
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"storyFile": *storyFile,
		}).Fatal("Failed to load story from file")
	}

	arcHandler, err := handlers.NewArcHandler(story)
	if err != nil {
		log.WithField("error", err).Fatal("Failed to create arc request handler")
	}

	router := mux.NewRouter()
	router.Handle("/{arc}", arcHandler)
	router.Handle("/", http.RedirectHandler("/intro", http.StatusFound))

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
	err = srv.Shutdown(ctx)
	if err != nil {
		log.WithField("error", err).Fatal("Error while attempting graceful shutdown")
	} else {
		log.Info("Server shut down gracefully")
	}
}
