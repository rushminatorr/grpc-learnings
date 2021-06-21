package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rushminator/carz/handlers"
	"go.uber.org/zap"
)

func main() {
	l, _ := zap.NewProduction()
	defer l.Sync() // flushes buffer, if any

	ch := handlers.NewCars(l)

	sm := http.NewServeMux()
	sm.Handle("/", ch)

	s := http.Server{
		Addr:    ":8989", // configure the bind address
		Handler: sm,      // set the default handler
		// ErrorLog:     logger,            // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Info("Starting server on port 8989")

		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			// l.Error("Error starting server: %s\n", zap.Error(err))
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
