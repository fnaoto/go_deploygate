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
	Error   bool                         `mapstructure:"error"`
	Results *GetAppMembersResponseResult `mapstructure:"results"`
}

type GetAppMembersResponseResult struct {
	Usage *Usage  `mapstructure:"usage"`
	Users []*User `mapstructure:"users"`
}

type Usage struct {
	Used uint `mapstructure:"used"`
	Max  uint `mapstructure:"max"`
}

type AddAppMembersResponse struct {
	Error   bool                         `mapstructure:"error"`
	Message string                       `mapstructure:"message"`
	Because string                       `mapstructure:"because"`
	Results *AddAppMembersResponseResult `mapstructure:"results"`
}

type AddAppMembersResponseResult struct {
	Invite  string  `mapstructure:"invite"`
	Added   []*User `mapstructure:"added"`
	Invited []*User `mapstructure:"invited"`
}

type RemoveAppMembersResponse struct {
	Error   bool                            `mapstructure:"error"`
	Message string                          `mapstructure:"message"`
	Because string                          `mapstructure:"because"`
	Results *RemoveAppMembersResponseResult `mapstructure:"results"`
}

type RemoveAppMembersResponseResult struct {
	Invite string `mapstructure:"invite"`
}
