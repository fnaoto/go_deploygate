package deploygate

// Request

type ListOrganizationTeamMembersRequest struct {
	Organization string
	Team         string
}

type AddOrganizationTeamMemberRequest struct {
	Organization string
	Team         string
	User         string
}

type RemoveOrganizationTeamMemberRequest struct {
	Organization string
	Team         string
	User         string
}

// Response

type ListOrganizationTeamMembersResponse struct {
	Error   bool     `json:"error"`
	Message string   `json:"message"`
	Users   []Member `json:"users"`
}

type AddOrganizationTeamMemberResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type RemoveOrganizationTeamMemberResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
