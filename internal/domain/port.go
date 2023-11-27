package domain

// PortData represents the data structure for port information.
type PortData struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

// Port represents the overall structure of a port, including its ID and data.
type Port struct {
	ID       string `json:"-"`
	PortData `json:""`
}
