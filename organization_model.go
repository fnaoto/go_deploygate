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
	Error         bool           `json:"error"`
	Message       string         `json:"message"`
	Organizations []Organization `json:"organizations"`
}

type CreateOrganizationResponse struct {
	Error        bool         `json:"error"`
	Message      string       `json:"message"`
	Organization Organization `json:"organization"`
}

type GetOrganizationResponse struct {
	Error        bool             `json:"error"`
	Message      string           `json:"message"`
	Organization OrganizationInfo `json:"organization"`
}

type UpdateOrganizationResponse struct {
	Error        bool         `json:"error"`
	Message      string       `json:"message"`
	Organization Organization `json:"organization"`
}

type DeleteOrganizationResponse struct {
	Error   bool    `json:"error"`
	Message string  `json:"message"`
	Results Results `json:"results"`
}
