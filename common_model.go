package deploygate

type Results struct {
	Message string `mapstructure:"message"`
}

type Organization struct {
	Type        string `mapstructure:"type"`
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Url         string `mapstructure:"url"`
}

type User struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type Member struct {
	Type    string `mapstructure:"type"`
	Name    string `mapstructure:"name"`
	IconUrl string `mapstructure:"icon_url"`
	Url     string `mapstructure:"url"`
}

type Teams struct {
	Name        string `mapstructure:"name"`
	Role        string `mapstructure:"role"`
	MemberCount uint   `mapstructure:"member_count"`
}
