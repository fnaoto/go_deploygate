package deploygate

type AppConfig struct {
	Owner    string
	Platform string
	AppId    string
	File     string
}

type UploadAppResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}
