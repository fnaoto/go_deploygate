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
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Results string `mapstructure:"results" json:"results"`
}

type DisableAppProtectionResponse struct {
	Error   bool   `mapstructure:"error" json:"error"`
	Message string `mapstructure:"message" json:"message"`
	Results string `mapstructure:"results" json:"results"`
}
