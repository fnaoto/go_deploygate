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
	Error   bool               `mapstructure:"error" json:"error"`
	Message string             `mapstructure:"message" json:"message"`
	Because string             `mapstructure:"because" json:"because"`
	Users   []EnterpriseMember `mapstructure:"users" json:"users"`
}

type AddEnterpriseMemberResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
}

type RemoveEnterpriseMemberResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
}
