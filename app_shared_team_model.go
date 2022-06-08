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
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
	Teams   []Team `json:"teams"`
}

type AddAppSharedTeamResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
}

type RemoveAppSharedTeamResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
}
