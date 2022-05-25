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

type DeleteAppMembersRequest struct {
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
	Usage *Usage    `mapstructure:"usage"`
	Users []*Member `mapstructure:"users"`
}

type Usage struct {
	Used uint `mapstructure:"used"`
	Max  uint `mapstructure:"max"`
}

type Member struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type AddAppMembersResponse struct {
	Error   bool                         `mapstructure:"error"`
	Message string                       `mapstructure:"message"`
	Because string                       `mapstructure:"because"`
	Results *AddAppMembersResponseResult `mapstructure:"results"`
}

type AddAppMembersResponseResult struct {
	Invite  string     `mapstructure:"invite"`
	Added   []*Added   `mapstructure:"added"`
	Invited []*Invited `mapstructure:"invited"`
}

type Added struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type Invited struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type DeleteAppMembersResponse struct {
	Error   bool                            `mapstructure:"error"`
	Message string                          `mapstructure:"message"`
	Because string                          `mapstructure:"because"`
	Results *DeleteAppMembersResponseResult `mapstructure:"results"`
}

type DeleteAppMembersResponseResult struct {
	Invite string `mapstructure:"invite"`
}
