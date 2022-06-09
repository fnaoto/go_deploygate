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
	Error   bool    `mapstructure:"error" json:"error"`
	Message string  `mapstructure:"message" json:"message"`
	Results Results `mapstructure:"results" json:"results"`
}
