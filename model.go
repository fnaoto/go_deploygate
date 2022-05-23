package deploygate

import (
	"io"
	"net/http"
	"net/url"
)

type HttpRequest struct {
	method string
	spath  string
	body   io.Reader
	header *Header
}

type Header struct {
	accept      string `default:"application/json"`
	contentType string `default:"application/x-www-form-urlencoded"`
}

type Client struct {
	endpoint   *url.URL     // URL Parsed endpoint
	apiKey     string       // DeployGate API Key
	httpClient *http.Client // Http client for using http request
}

type AppMemberConfig struct {
	Owner    string
	Platform string
	AppId    string
	Users    string
	Role     string
	File     string
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

type AddAppMemberResponse struct {
	Error   bool                        `mapstructure:"error"`
	Message string                      `mapstructure:"message"`
	Because string                      `mapstructure:"because"`
	Results *AddAppMemberResponseResult `mapstructure:"results"`
}

type AddAppMemberResponseResult struct {
	Invite  string     `mapstructure:"invite"`
	Added   []*Added   `mapstructure:"added"`
	Invited []*Invited `mapstructure:"invited"`
}

type Added struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type Invited struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type DeleteAppMemberResponse struct {
	Error   bool                           `mapstructure:"error"`
	Message string                         `mapstructure:"message"`
	Because string                         `mapstructure:"because"`
	Results *DeleteAppMemberResponseResult `mapstructure:"results"`
}

type DeleteAppMemberResponseResult struct {
	Invite string `mapstructure:"invite"`
}

type UploadAppResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}
