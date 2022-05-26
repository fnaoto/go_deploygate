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
