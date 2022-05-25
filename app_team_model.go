package deploygate

// Request

type AppTeamConfig struct {
	Organizations string
	Platform      string
	AppId         string
	Team          string
}

// Response

type GetAppTeamsResponse struct {
	Error bool     `mapstructure:"error"`
	Teams []*Teams `mapstructure:"teams"`
}

type Teams struct {
	Name        string `mapstructure:"name"`
	Role        string `mapstructure:"role"`
	MemberCount uint   `mapstructure:"member_count"`
}

type AddAppTeamsResponse struct {
	Error bool     `mapstructure:"error"`
	Teams []*Teams `mapstructure:"teams"`
}
