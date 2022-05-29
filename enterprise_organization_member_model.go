package deploygate

// Request

type ListEnterpriseOrganizationMembersRequest struct {
	Enterprise   string
	Organization string
}

type AddEnterpriseOrganizationMembersRequest struct {
	Enterprise   string
	Organization string
	User         string
}

type RemoveEnterpriseOrganizationMembersRequest struct {
	Enterprise   string
	Organization string
	User         string
}

// Response

type ListEnterpriseOrganizationMembersResponse struct {
	Error   bool    `mapstructures:"error"`
	Message string  `mapstructure:"message"`
	Because string  `mapstructure:"because"`
	Users   []*User `mapstructures:"users"`
}

type AddEnterpriseOrganizationMembersResponse struct {
	Error   bool   `mapstructures:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}

type RemoveEnterpriseOrganizationMembersResponse struct {
	Error   bool   `mapstructures:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}
