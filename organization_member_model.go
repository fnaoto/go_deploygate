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
	Error   bool   `mapstructures:"error"`
	Message string `mapstructures:"message"`
	Members []*Member
}

type AddOrganizationMemberResponse struct {
	Error   bool   `mapstructures:"error"`
	Message string `mapstructures:"message"`
}

type RemoveOrganizationMemberResponse struct {
	Error   bool   `mapstructures:"error"`
	Message string `mapstructures:"message"`
}
