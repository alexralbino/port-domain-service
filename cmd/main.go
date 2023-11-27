package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexralbino/port-domain-service/internal/repository"
	"github.com/alexralbino/port-domain-service/internal/service"
)

func main() {
	// Initialize the in-memory repository
	repo := repository.NewInMemoryRepository()

	// Initialize the port service with the repository
	portService := service.NewPortService(repo)

	// Handle signals for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-stop
		log.Printf("Received signal %v. Shutting down gracefully...", sig)
		// Perform any cleanup or additional shutdown logic here
		os.Exit(0)
	}()

	// Import ports.json file
	err := portService.LoadPortsFromFile("import-files/ports.json")
	if err != nil {
		log.Println(err)
	}

	// testing retrieving a record
	log.Println(repo.GetByID("JPTOY"))

	log.Println("Port Service is running. Press Ctrl+C to stop.")
	select {}
}
