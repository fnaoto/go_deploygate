package deploygate

// Request

type ListAppTeamsRequest struct {
	Organization string
	Platform     string
	AppId        string
}

type AddAppTeamRequest struct {
	Organization string
	Platform     string
	AppId        string
	Team         string
}

type RemoveAppTeamRequest struct {
	Organization string
	Platform     string
	AppId        string
	Team         string
}

// Response

type ListAppTeamsResponse struct {
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

type AddAppTeamResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}

type RemoveAppTeamResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}
