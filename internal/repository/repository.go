package repository

import "github.com/alexralbino/port-domain-service/internal/domain"

// PortRepository defines the interface for interacting with port data.
type PortRepository interface {
	Save(port *domain.Port)
	GetByID(id string) (*domain.Port, error)
}

// NewInMemoryRepository creates a new in-memory repository.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		ports: make(map[string]*domain.Port),
	}
}
