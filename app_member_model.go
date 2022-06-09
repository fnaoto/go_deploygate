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
	Error   bool                        `mapstructure:"error" json:"error"`
	Results GetAppMembersResponseResult `mapstructure:"results" json:"results"`
}

type GetAppMembersResponseResult struct {
	Usage Usage  `mapstructure:"usage" json:"usage"`
	Users []User `mapstructure:"users" json:"users"`
}

type Usage struct {
	Used uint `mapstructure:"used" json:"used"`
	Max  uint `mapstructure:"max" json:"max"`
}

type AddAppMembersResponse struct {
	Error   bool                        `mapstructure:"error" json:"error"`
	Message string                      `mapstructure:"message" json:"message"`
	Because string                      `mapstructure:"because" json:"because"`
	Results AddAppMembersResponseResult `mapstructure:"results" json:"results"`
}

type AddAppMembersResponseResult struct {
	Invite  string `mapstructure:"invite" json:"invite"`
	Added   []User `mapstructure:"added" json:"added"`
	Invited []User `mapstructure:"invited" json:"invited"`
}

type RemoveAppMembersResponse struct {
	Error   bool                           `mapstructure:"error" json:"error"`
	Message string                         `mapstructure:"message" json:"message"`
	Because string                         `mapstructure:"because" json:"because"`
	Results RemoveAppMembersResponseResult `mapstructure:"results" json:"results"`
}

type RemoveAppMembersResponseResult struct {
	Invite string `mapstructure:"invite" json:"invite"`
}
