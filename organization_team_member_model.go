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
	Error   bool     `mapstructure:"error" json:"error"`
	Message string   `mapstructure:"message" json:"message"`
	Users   []Member `mapstructure:"users" json:"users"`
}

type AddOrganizationTeamMemberResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
}

type RemoveOrganizationTeamMemberResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
}
