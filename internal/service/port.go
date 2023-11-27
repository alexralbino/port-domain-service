package service

import (
	"encoding/json"
	"io"
	"os"

	"errors"

	"github.com/alexralbino/port-domain-service/internal/domain"
	"github.com/alexralbino/port-domain-service/internal/repository"
)

// PortService represents a service for handling port-related operations.
type PortService struct {
	repo repository.PortRepository
}

// NewPortService creates a new instance of PortService with the provided repository.
func NewPortService(repo repository.PortRepository) *PortService {
	return &PortService{
		repo: repo,
	}
}

// LoadPortsFromFile loads port data from a JSON file and saves it to the repository.
func (s *PortService) LoadPortsFromFile(filePath string) error {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("failed to open file")
	}
	defer file.Close() // Ensure the file is closed after the function returns

	// Create a JSON decoder for reading the file
	decoder := json.NewDecoder(file)

	// Read the opening brace of the JSON object
	_, err = decoder.Token()
	if err != nil {
		return err
	}

	// Loop through the JSON tokens
	for {
		// Read the next token (port ID) from the JSON stream
		id, err := decoder.Token()
		if err != nil {
			// End of JSON stream
			return err
		}

		// Create a new Port instance
		port := &domain.Port{}

		// Decode the JSON object representing the port
		err = decoder.Decode(&port)
		if err == io.EOF {
			// End of file
			break
		}

		if err != nil {
			// Error decoding JSON object
			return err
		}

		// Set the port ID
		port.ID = id.(string)

		// Validate the port data
		if err := port.IsValid(); err != nil {
			return err
		}

		// Save the port to the repository
		s.repo.Save(port)
	}

	// Successfully loaded and processed all ports from the file
	return nil
}
