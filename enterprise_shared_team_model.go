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
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
	Teams   []Team `json:"team"`
}

type AddEnterpriseSharedTeamResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
}

type RemoveEnterpriseSharedTeamResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
}
