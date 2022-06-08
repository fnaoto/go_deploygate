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
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
	Teams   []Team `json:"teams"`
}

type AddAppTeamResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
}

type RemoveAppTeamResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
}
