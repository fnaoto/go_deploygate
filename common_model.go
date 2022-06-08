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

type OrganizationInfo struct {
	Type         string        `json:"type"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Url          string        `json:"url"`
	Teams        []Team        `json:"teams"`
	Applications []Application `json:"applications"`
	Members      []Member      `json:"members"`
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

type EnterpriseMember struct {
	Type         string `json:"type"`
	Name         string `json:"name"`
	IconUrl      string `json:"icon_url"`
	Url          string `json:"url"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	CreatedAt    string `json:"created_at"`
	LastAccessAt string `json:"last_access_at"`
}

type Team struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	MemberCount uint   `json:"member_count"`
}

type Application struct {
	Name             string      `json:"name"`
	PackageName      string      `json:"package_name"`
	Labels           interface{} `json:"labels"`
	OsName           string      `json:"os_name"`
	Path             string      `json:"path"`
	UpdatedAt        uint        `json:"updated_at"`
	VersionCode      string      `json:"version_code"`
	VersionName      string      `json:"version_name"`
	SdkVersion       uint        `json:"sdk_version"`
	RawSdkVersion    string      `json:"raw_sdk_version"`
	TargetSdkVersion uint        `json:"target_sdk_version"`
	Signature        string      `json:"signature"`
	Md5              string      `json:"md5"`
	Revision         uint        `json:"revision"`
	File_Size        string      `json:"file_size"`
	Icon             string      `json:"icon"`
	Message          string      `json:"message"`
	File             string      `json:"file"`
	User             struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		ProfileIcon string `json:"profile_icon"`
	} `json:"user"`
}
