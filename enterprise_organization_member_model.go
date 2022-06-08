package deploygate

// Request

type ListEnterpriseOrganizationMembersRequest struct {
	Enterprise   string
	Organization string
}

type AddEnterpriseOrganizationMemberRequest struct {
	Enterprise   string
	Organization string
	User         string
}

type RemoveEnterpriseOrganizationMemberRequest struct {
	Enterprise   string
	Organization string
	User         string
}

// Response

type ListEnterpriseOrganizationMembersResponse struct {
	Error   bool     `json:"error"`
	Message string   `json:"message"`
	Because string   `json:"because"`
	Users   []Member `json:"users"`
}

type AddEnterpriseOrganizationMemberResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
}

type RemoveEnterpriseOrganizationMemberResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
}
