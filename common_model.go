package deploygate

type Results struct {
	Message string `mapstructure:"message" json:"message"`
}

type Organization struct {
	Type        string `mapstructure:"type" json:"type"`
	Name        string `mapstructure:"name" json:"name"`
	Description string `mapstructure:"description" json:"description"`
	Url         string `mapstructure:"url" json:"url"`
}

type OrganizationInfo struct {
	Type         string        `mapstructure:"type" json:"type"`
	Name         string        `mapstructure:"name" json:"name"`
	Description  string        `mapstructure:"description" json:"description"`
	Url          string        `mapstructure:"url" json:"url"`
	Teams        []Team        `mapstructure:"teams" json:"teams"`
	Applications []Application `mapstructure:"applications" json:"applications"`
	Members      []Member      `mapstructure:"members" json:"members"`
}

type User struct {
	Name string `mapstructure:"name" json:"name"`
	Role uint   `mapstructure:"role" json:"role"`
}

type Member struct {
	Type    string `mapstructure:"type" json:"type"`
	Name    string `mapstructure:"name" json:"name"`
	IconUrl string `mapstructure:"icon_url" json:"icon_url"`
	Url     string `mapstructure:"url" json:"url"`
}

type EnterpriseMember struct {
	Type         string `mapstructure:"type" json:"type"`
	Name         string `mapstructure:"name" json:"name"`
	IconUrl      string `mapstructure:"icon_url" json:"icon_url"`
	Url          string `mapstructure:"url" json:"url"`
	FullName     string `mapstructure:"full_name" json:"full_name"`
	Email        string `mapstructure:"email" json:"email"`
	Role         string `mapstructure:"role" json:"role"`
	CreatedAt    string `mapstructure:"created_at" json:"created_at"`
	LastAccessAt string `mapstructure:"last_access_at" json:"last_access_at"`
}

type Team struct {
	Name        string `mapstructure:"name" json:"name"`
	Role        string `mapstructure:"role" json:"role"`
	MemberCount uint   `mapstructure:"member_count" json:"member_count"`
}

type Application struct {
	Name             string      `mapstructure:"name" json:"name"`
	PackageName      string      `mapstructure:"package_name" json:"package_name"`
	Labels           interface{} `mapstructure:"labels" json:"labels"`
	OsName           string      `mapstructure:"os_name" json:"os_name"`
	Path             string      `mapstructure:"path" json:"path"`
	UpdatedAt        uint        `mapstructure:"updated_at" json:"updated_at"`
	VersionCode      string      `mapstructure:"version_code" json:"version_code"`
	VersionName      string      `mapstructure:"version_name" json:"version_name"`
	SdkVersion       uint        `mapstructure:"sdk_version" json:"sdk_version"`
	RawSdkVersion    string      `mapstructure:"raw_sdk_version" json:"raw_sdk_version"`
	TargetSdkVersion uint        `mapstructure:"target_sdk_version" json:"target_sdk_version"`
	Signature        string      `mapstructure:"signature" json:"signature"`
	Md5              string      `mapstructure:"md5" json:"md5"`
	Revision         uint        `mapstructure:"revision" json:"revision"`
	File_Size        string      `mapstructure:"file_size" json:"file_size"`
	Icon             string      `mapstructure:"icon" json:"icon"`
	Message          string      `mapstructure:"message" json:"message"`
	File             string      `mapstructure:"file" json:"file"`
	User             struct {
		ID          uint   `mapstructure:"id" json:"id"`
		Name        string `mapstructure:"name" json:"name"`
		ProfileIcon string `mapstructure:"profile_icon" json:"profile_icon"`
	} `mapstructure:"user" json:"user"`
}
