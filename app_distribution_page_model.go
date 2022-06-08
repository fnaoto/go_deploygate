package deploygate

// Request

type DeleteAppDistributionsPageRequest struct {
	Owner            string
	Platform         string
	AppId            string
	DistributionName string
}

// Response

type DeleteAppDistributionsPageResponse struct {
	Error   bool    `json:"error"`
	Message string  `json:"message"`
	Results Results `json:"results"`
}
