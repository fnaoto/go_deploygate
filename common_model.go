package deploygate

type Results struct {
	Message string `mapstructure:"message"`
}

type Organization struct {
	Type        string `mapstructures:"type"`
	Name        string `mapstructures:"name"`
	Description string `mapstructures:"description"`
	Url         string `mapstructures:"url"`
}

type Member struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type User struct {
	Type    string `mapstructure:"type"`
	Name    string `mapstructure:"name"`
	IconUrl string `mapstructure:"iconurl"`
	Url     string `mapstructure:"url"`
}

type Teams struct {
	Name        string `mapstructure:"name"`
	Role        string `mapstructure:"role"`
	MemberCount uint   `mapstructure:"member_count"`
}
