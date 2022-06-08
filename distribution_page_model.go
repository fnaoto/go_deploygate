package deploygate

// Request

type DeleteDistributionsPageRequest struct {
	Distribution string
}

// Response

type DeleteDistributionsPageResponse struct {
	Error   bool    `json:"error"`
	Message string  `json:"message"`
	Results Results `json:"results"`
}
