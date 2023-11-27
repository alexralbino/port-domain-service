package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize the in-memory repository
	// repo := repository.NewInMemoryRepository()

	// Handle signals for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-stop
		log.Printf("Received signal %v. Shutting down gracefully...", sig)
		// Perform any cleanup or additional shutdown logic here
		os.Exit(0)
	}()

	log.Println("Port Service is running. Press Ctrl+C to stop.")
	select {}
}
