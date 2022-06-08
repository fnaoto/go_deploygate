package deploygate

type Results struct {
	Message string `json:"message"`
}

type Organization struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type User struct {
	Name string `json:"name"`
	Role uint   `json:"role"`
}

type Member struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	IconUrl string `json:"icon_url"`
	Url     string `json:"url"`
}

type Team struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	MemberCount uint   `json:"member_count"`
}
