package deploygate

// Request

// Response

type ListOrganizationsResponse struct {
	Error         bool             `mapstructures:"error"`
	Message       string           `mapstructures:"message"`
	Organizations []*Organizations `mapstructures:"organizations"`
}

type Organizations struct {
	Type        string `mapstructures:"type"`
	Name        string `mapstructures:"name"`
	Description string `mapstructures:"description"`
	Url         string `mapstructures:"url"`
}
