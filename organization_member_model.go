package deploygate

// Request

type ListOrganizationMembersRequest struct {
	Organization string
}

// Response

type ListOrganizationMembersResponse struct {
	Error   bool   `mapstructures:"error"`
	Message string `mapstructures:"message"`
	Members []*Member
}
