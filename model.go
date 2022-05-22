package deploygate

import (
	"net/http"
	"net/url"
)

type Client struct {
	endpoint   *url.URL     // URL Parsed endpoint
	apiKey     string       // DeployGate API Key
	httpClient *http.Client // Http client for using http request
}

type AppMemberConfig struct {
	Owner    string
	Platform string
	AppId    string
}

type Member struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type GetAppMemberResponse struct {
	Error   bool                        `mapstructure:"error"`
	Results *GetAppMemberResponseResult `mapstructure:"results"`
}

type GetAppMemberResponseResult struct {
	Usage *Usage    `mapstructure:"usage"`
	Users []*Member `mapstructure:"users"`
}

type Usage struct {
	Used uint `mapstructure:"used"`
	Max  uint `mapstructure:"max"`
}
