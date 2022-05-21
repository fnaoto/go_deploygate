package deploygate

import (
	"net/http"
	"net/url"
)

type Client struct {
	// URL Parsed endpoint
	endpoint *url.URL

	// DeployGate API Key
	apiKey string

	// Http client for using http request
	httpClient *http.Client
}
