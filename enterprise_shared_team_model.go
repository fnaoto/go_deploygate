package deploygate

// Request

type ListEnterpriseSharedTeamsRequest struct {
	Enterprise string
}

type AddEnterpriseSharedTeamRequest struct {
	Enterprise  string
	SharedTeam  string
	Description string
}

type RemoveEnterpriseSharedTeamRequest struct {
	Enterprise string
	SharedTeam string
}

// Response

type ListEnterpriseSharedTeamsResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
	Teams   []Team `mapstructure:"teams" json:"teams"`
}

type AddEnterpriseSharedTeamResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
}

type RemoveEnterpriseSharedTeamResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Because string `mapstructure:"because" json:"because"`
}
