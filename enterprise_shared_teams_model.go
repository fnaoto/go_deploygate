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
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
	Teams   []*Teams
}

type AddEnterpriseSharedTeamResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}

type RemoveEnterpriseSharedTeamResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}
