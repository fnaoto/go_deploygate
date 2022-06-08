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
	Error         bool            `mapstructure:"error"`
	Message       string          `mapstructure:"message"`
	Organizations []*Organization `mapstructure:"organizations"`
}

type CreateOrganizationResponse struct {
	Error        bool          `mapstructure:"error"`
	Message      string        `mapstructure:"message"`
	Organization *Organization `mapstructure:"organization"`
}

type GetOrganizationResponse struct {
	Error        bool          `mapstructure:"error"`
	Message      string        `mapstructure:"message"`
	Organization *Organization `mapstructure:"organization"`
}

type UpdateOrganizationResponse struct {
	Error        bool          `mapstructure:"error"`
	Message      string        `mapstructure:"message"`
	Organization *Organization `mapstructure:"organization"`
}

type DeleteOrganizationResponse struct {
	Error   bool     `mapstructure:"error"`
	Message string   `mapstructure:"message"`
	Results *Results `mapstructure:"results"`
}
