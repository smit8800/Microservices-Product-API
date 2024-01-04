package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/handlers"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gg := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gg)

	s := &http.Server{
		Addr:         ":8000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Start the HTTP server in a new goroutine to run concurrently.
	go func() {
		// Start listening for incoming HTTP requests.
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Create a channel to receive OS signals like interrupts or termination signals.
	sigChan := make(chan os.Signal)

	// Notify the sigChan channel when the program receives the OS interrupt signal or the OS kill.
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block the program until a signal is received on the sigChan channel.
	// Log a message indicating that a signal was received and specify which signal was received.
	sig := <-sigChan
	l.Println("received terminate, graceful shutdown", sig)

	// Create a context with a 30-second timeout for the server shutdown operation.
	// Gracefully shut down the HTTP server using the created context.
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
