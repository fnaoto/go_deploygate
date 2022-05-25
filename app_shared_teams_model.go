package deploygate

// Request

type ListAppSharedTeamsRequest struct {
	Organizations string
	Platform      string
	AppId         string
}

// Response

type ListAppSharedTeamsResponse struct {
	Error bool `mapstructure:"error"`
	Teams []*Teams
}
