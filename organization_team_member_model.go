package deploygate

// Request

type ListOrganizationTeamMembersRequest struct {
	Organization string
	Team         string
}

// Response

type ListOrganizationTeamMembersResponse struct {
	Error bool    `mapstructure:"error"`
	Users []*User `mapstructure:"users"`
}

type User struct {
	Type    string `mapstructure:"type"`
	Name    string `mapstructure:"name"`
	IconUrl string `mapstructure:"iconurl"`
	Url     string `mapstructure:"url"`
}
