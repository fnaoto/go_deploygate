package deploygate

// Request

type ListAppTeamsRequest struct {
	Organization string
	Platform     string
	AppId        string
}

type AddAppTeamsRequest struct {
	Organization string
	Platform     string
	AppId        string
	Team         string
}

type RemoveAppTeamsRequest struct {
	Organization string
	Platform     string
	AppId        string
	Team         string
}

// Response

type ListAppTeamsResponse struct {
	Error   bool     `mapstructure:"error"`
	Message string   `mapstructure:"message"`
	Because string   `mapstructure:"because"`
	Teams   []*Teams `mapstructure:"teams"`
}

type Teams struct {
	Name        string `mapstructure:"name"`
	Role        string `mapstructure:"role"`
	MemberCount uint   `mapstructure:"member_count"`
}

type AddAppTeamsResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}

type RemoveAppTeamsResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}
