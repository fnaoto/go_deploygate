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
	Error   bool     `mapstructure:"error"`
	Message string   `mapstructure:"message"`
	Results *Results `mapstructure:"results"`
}
