package deploygate

// Request

type UploadAppsRequest struct {
	Owner    string
	Platform string
	AppId    string
	File     string
}

// Response

type UploadAppsResponse struct {
	Error   bool        `mapstructure:"error" json:"error"`
	Message string      `mapstructure:"message" json:"message"`
	Because string      `mapstructure:"because" json:"because"`
	Results Application `mapstructure:"results" json:"results"`
}
