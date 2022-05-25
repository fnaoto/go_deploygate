package deploygate

// Request

type ListAppSharedTeamsRequest struct {
	Organizations string
	Platform      string
	AppId         string
}

type AddAppSharedTeamsRequest struct {
	Organizations string
	Platform      string
	AppId         string
	Team          string
}

type RemoveAppSharedTeamsRequest struct {
	Organizations string
	Platform      string
	AppId         string
	Team          string
}

// Response

type ListAppSharedTeamsResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
	Teams   []*Teams
}

type AddAppSharedTeamsResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}

type RemoveAppSharedTeamsResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}
