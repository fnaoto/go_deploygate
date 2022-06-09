package deploygate

// Request

type ListAppSharedTeamsRequest struct {
	Organization string
	Platform     string
	AppId        string
}

type AddAppSharedTeamRequest struct {
	Organization string
	Platform     string
	AppId        string
	Team         string
}

type RemoveAppSharedTeamRequest struct {
	Organization string
	Platform     string
	AppId        string
	Team         string
}

// Response

type ListAppSharedTeamsResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
	Teams   []Team `mapstructure:"teams" json:"teams"`
}

type AddAppSharedTeamResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
}

type RemoveAppSharedTeamResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
}
