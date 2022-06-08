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
	Error   bool               `json:"error"`
	Message string             `json:"message"`
	Because string             `json:"because"`
	Users   []EnterpriseMember `json:"users"`
}

type AddEnterpriseMemberResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
}

type RemoveEnterpriseMemberResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
}
