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
	accept      string
	contentType string
}

type Client struct {
	endpoint   *url.URL     // URL Parsed endpoint
	apiKey     string       // DeployGate API Key
	httpClient *http.Client // Http client for using http request
}
