package deploygate

// Request

type EnableAppProtectionRequest struct {
	Owner    string
	Platform string
	AppId    string
	Revision int32
}

type DisableAppProtectionRequest struct {
	Owner    string
	Platform string
	AppId    string
	Revision int32
}

// Response

type EnableAppProtectionResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Results string `json:"results"`
}

type DisableAppProtectionResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Results string `json:"results"`
}
