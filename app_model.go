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
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}
