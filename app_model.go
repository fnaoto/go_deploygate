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
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Because string `json:"because"`
}
