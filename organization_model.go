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
	Error         bool           `mapstructure:"error" json:"error"`
	Message       string         `mapstructure:"message" json:"message"`
	Organizations []Organization `mapstructure:"organizations" json:"organizations"`
}

type CreateOrganizationResponse struct {
	Error        bool         `mapstructure:"error" json:"error"`
	Message      string       `mapstructure:"message" json:"message"`
	Organization Organization `mapstructure:"organization" json:"organization"`
}

type GetOrganizationResponse struct {
	Error        bool             `mapstructure:"error" json:"error"`
	Message      string           `mapstructure:"message" json:"message"`
	Organization OrganizationInfo `mapstructure:"organization" json:"organization"`
}

type UpdateOrganizationResponse struct {
	Error        bool         `mapstructure:"error" json:"error"`
	Message      string       `mapstructure:"message" json:"message"`
	Organization Organization `mapstructure:"organization" json:"organization"`
}

type DeleteOrganizationResponse struct {
	Error   bool    `mapstructure:"error" json:"error"`
	Message string  `mapstructure:"message" json:"message"`
	Results Results `mapstructure:"results" json:"results"`
}
