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

// Response

type ListOrganizationTeamMembersResponse struct {
	Error   bool    `mapstructure:"error"`
	Message string  `mapstructure:"message"`
	Users   []*User `mapstructure:"users"`
}

type User struct {
	Type    string `mapstructure:"type"`
	Name    string `mapstructure:"name"`
	IconUrl string `mapstructure:"iconurl"`
	Url     string `mapstructure:"url"`
}

type AddOrganizationTeamMemberResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
}
