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
	Error   bool     `mapstructure:"error" json:"error"`
	Message string   `mapstructure:"message" json:"message"`
	Because string   `mapstructure:"because" json:"because"`
	Users   []Member `mapstructure:"users" json:"users"`
}

type AddEnterpriseOrganizationMemberResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
}

type RemoveEnterpriseOrganizationMemberResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
}
