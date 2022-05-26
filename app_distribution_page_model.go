package deploygate

// Request

type DeleteAppDistributionsPageRequest struct {
	Distribution string
}

// Response

type DeleteAppDistributionsPageResponse struct {
	Error   bool     `mapstructure:"error"`
	Message string   `mapstructure:"message"`
	Results *Results `mapstructure:"results"`
}

type Results struct {
	Message string `mapstructure:"message"`
}
