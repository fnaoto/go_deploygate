package deploygate

// Request

type ListOrganizationMembersRequest struct {
	Organization string
}

type AddOrganizationMemberByUserNameRequest struct {
	Organization string
	UserName     string
}

type AddOrganizationMemberByEmailRequest struct {
	Organization string
	Email        string
}

type RemoveOrganizationMemberByUserNameRequest struct {
	Organization string
	UserName     string
}

type RemoveOrganizationMemberByEmailRequest struct {
	Organization string
	Email        string
}

// Response

type ListOrganizationMembersResponse struct {
	Error   bool     `mapstructure:"error" json:"error"`
	Message string   `mapstructure:"message" json:"message"`
	Members []Member `mapstructure:"members" json:"members"`
}

type AddOrganizationMemberResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
}

type RemoveOrganizationMemberResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
}
