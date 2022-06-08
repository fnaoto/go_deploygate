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
	Error   bool      `mapstructure:"error"`
	Message string    `mapstructure:"message"`
	Because string    `mapstructure:"because"`
	Users   []*Member `mapstructure:"users"`
}

type AddEnterpriseOrganizationMemberResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}

type RemoveEnterpriseOrganizationMemberResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}
