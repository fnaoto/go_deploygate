package deploygate

// Request

type DeleteDistributionsPageRequest struct {
	Distribution string
}

// Response

type DeleteDistributionsPageResponse struct {
	Error   bool     `mapstructure:"error"`
	Message string   `mapstructure:"message"`
	Results *Results `mapstructure:"results"`
}

type Results struct {
	Message string `mapstructure:"message"`
}
