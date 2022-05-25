package deploygate

// Request

type GetAppTeamsRequest struct {
	Organizations string
	Platform      string
	AppId         string
}

type AddAppTeamsRequest struct {
	Organizations string
	Platform      string
	AppId         string
	Team          string
}

type RemoveAppTeamsRequest struct {
	Organizations string
	Platform      string
	AppId         string
	Team          string
}

// Response

type GetAppTeamsResponse struct {
	Error   bool     `mapstructure:"error"`
	Message string   `mapstructure:"message"`
	Because string   `mapstructure:"because"`
	Teams   []*Teams `mapstructure:"teams"`
}

type Teams struct {
	Name        string `mapstructure:"name"`
	Role        string `mapstructure:"role"`
	MemberCount uint   `mapstructure:"member_count"`
}

type AddAppTeamsResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}

type RemoveAppTeamsResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}
