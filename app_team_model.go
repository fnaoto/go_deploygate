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
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
	Teams   []Team `mapstructure:"teams" json:"teams"`
}

type AddAppTeamResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
}

type RemoveAppTeamResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
}
