package deploygate

// Request

type EnableAppProtectionRequest struct {
	Owner    string
	Platform string
	AppId    string
	Revision int32
}

// Response

type EnableAppProtectionResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Results string `mapstructure:"results"`
}
