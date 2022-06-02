package deploygate

// Request

type CreateOrganizationRequest struct {
	Name        string
	Description string
}

type GetOrganizationRequest struct {
	Name string
}

type UpdateOrganizationRequest struct {
	Name        string
	Description string
}

type DeleteOrganizationRequest struct {
	Name string
}

// Response

type ListOrganizationsResponse struct {
	Error         bool            `mapstructures:"error"`
	Message       string          `mapstructures:"message"`
	Organizations []*Organization `mapstructures:"organizations"`
}

type CreateOrganizationResponse struct {
	Error        bool          `mapstructures:"error"`
	Message      string        `mapstructures:"message"`
	Organization *Organization `mapstructures:"organization"`
}

type GetOrganizationResponse struct {
	Error        bool          `mapstructures:"error"`
	Message      string        `mapstructures:"message"`
	Organization *Organization `mapstructures:"organization"`
}

type UpdateOrganizationResponse struct {
	Error        bool          `mapstructures:"error"`
	Message      string        `mapstructures:"message"`
	Organization *Organization `mapstructures:"organization"`
}

type DeleteOrganizationResponse struct {
	Error   bool     `mapstructures:"error"`
	Message string   `mapstructures:"message"`
	Results *Results `mapstructures:"results"`
}
