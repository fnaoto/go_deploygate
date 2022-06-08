package deploygate

// Request

type ListEnterpriseMembersRequest struct {
	Enterprise string
}

type AddEnterpriseMemberRequest struct {
	Enterprise string
	User       string
}

type RemoveEnterpriseMemberRequest struct {
	Enterprise string
	User       string
}

// Response

type ListEnterpriseMembersResponse struct {
	Error   bool      `mapstructure:"error"`
	Message string    `mapstructure:"message"`
	Because string    `mapstructure:"because"`
	Users   []*Member `mapstructure:"users"`
}

type AddEnterpriseMemberResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}

type RemoveEnterpriseMemberResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}
