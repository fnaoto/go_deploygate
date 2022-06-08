package deploygate

// Request

type GetAppMembersRequest struct {
	Owner    string
	Platform string
	AppId    string
}

type AddAppMembersRequest struct {
	Owner    string
	Platform string
	AppId    string
	Users    string
	Role     string
}

type RemoveAppMembersRequest struct {
	Owner    string
	Platform string
	AppId    string
	Users    string
}

// Response

type GetAppMembersResponse struct {
	Error   bool                        `json:"error"`
	Results GetAppMembersResponseResult `json:"results"`
}

type GetAppMembersResponseResult struct {
	Usage Usage  `json:"usage"`
	Users []User `json:"users"`
}

type Usage struct {
	Used uint `json:"used"`
	Max  uint `json:"max"`
}

type AddAppMembersResponse struct {
	Error   bool                        `json:"error"`
	Message string                      `json:"message"`
	Because string                      `json:"because"`
	Results AddAppMembersResponseResult `json:"results"`
}

type AddAppMembersResponseResult struct {
	Invite  string `json:"invite"`
	Added   []User `json:"added"`
	Invited []User `json:"invited"`
}

type RemoveAppMembersResponse struct {
	Error   bool                           `json:"error"`
	Message string                         `json:"message"`
	Because string                         `json:"because"`
	Results RemoveAppMembersResponseResult `json:"results"`
}

type RemoveAppMembersResponseResult struct {
	Invite string `json:"invite"`
}
