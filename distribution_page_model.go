package deploygate

// Request

type DeleteDistributionsPageRequest struct {
	Distribution string
}

// Response

type DeleteDistributionsPageResponse struct {
	Error   bool    `mapstructure:"error" json:"error"`
	Message string  `mapstructure:"message" json:"message"`
	Results Results `mapstructure:"results" json:"results"`
}
