package service

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/alexralbino/port-domain-service/internal/domain"
	"github.com/stretchr/testify/assert"
)

// MockPortRepository is a mock implementation of the PortRepository interface for testing.
type MockPortRepository struct {
	SavedPort *domain.Port
}

// Save saves a port in the mock repository.
func (m *MockPortRepository) Save(port *domain.Port) {
	m.SavedPort = port
}

// GetByID retrieves a port from the mock repository by ID.
func (m *MockPortRepository) GetByID(id string) (*domain.Port, error) {
	if m.SavedPort != nil && m.SavedPort.ID == id {
		return m.SavedPort, nil
	}
	return nil, errors.New("record not found")
}

func TestLoadPortsFromFile(t *testing.T) {
	// Get the temporary file directory
	tempDir := os.TempDir()

	// Join path with temporary directory
	testFilePath := filepath.Join(tempDir, "json-test-file.json")

	tests := []struct {
		name          string
		fileName      string
		input         []byte
		expectedError string
		expectsError  bool
	}{
		{
			name:     "valid file",
			fileName: "json-test-file.json",
			input: []byte(`{
                "testport":{
                    "name": "testName",
                    "city": "testCity",
                    "country": "testCountry",
                    "alias": [],
                    "regions": [],
                    "coordinates": [
                    55.5130433,
                    25.4050165
                    ],
                    "province": "testProvince",
                    "timezone": "test/Timezone",
                    "unlocs": [
                    "testUnlocs"
                    ],
                    "code": "testCode"
                }
            }`),
			expectedError: "",
			expectsError:  false,
		},
		{
			name:          "file not exists",
			fileName:      "file-not-exists.json",
			input:         []byte(`{}`),
			expectedError: "failed to open file",
			expectsError:  true,
		},
		{
			name:          "invalid json format",
			fileName:      "json-test-file.json",
			input:         []byte(`invalid json`),
			expectedError: "failed to retrieve id token",
			expectsError:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := os.WriteFile(testFilePath, test.input, 0644)
			assert.NoError(t, err)
			defer os.Remove(testFilePath)

			// Create a mock repository for testing
			mockRepo := &MockPortRepository{}

			// Create a port service with the mock repository
			portService := NewPortService(mockRepo)

			loadFilePath := filepath.Join(tempDir, test.fileName)

			// Test LoadPortsFromFile using the temporary test file
			err = portService.LoadPortsFromFile(loadFilePath)

			// Checks if error is expected
			if test.expectsError {
				assert.Error(t, err, test.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
