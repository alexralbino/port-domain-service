package repository

import (
	"errors"
	"sync"

	"github.com/alexralbino/port-domain-service/internal/domain"
)

// InMemoryRepository represents an in-memory repository for storing port data.
type InMemoryRepository struct {
	mu    sync.RWMutex            // Mutex for concurrent access to the in-memory data
	ports map[string]*domain.Port // Map to store port data with port ID as key
}

// Save saves the given port to the in-memory repository.
func (r *InMemoryRepository) Save(port *domain.Port) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Store the port in the map with its ID as the key
	r.ports[port.ID] = port
}

// GetByID retrieves a port from the in-memory repository based on its ID.
func (r *InMemoryRepository) GetByID(id string) (*domain.Port, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Retrieve the port from the map
	port, exists := r.ports[id]
	if !exists {
		return nil, errors.New("record not found")
	}

	return port, nil
}
