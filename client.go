package deploygate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

const (
	DGApiTokenEnv = "DEPLOYGATE_API_KEY"
	DGApiEndpoint = "https://deploygate.com/api"
)

func DefaultClient() (*Client, error) {
	c, err := NewClient(os.Getenv(DGApiTokenEnv), DGApiEndpoint)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewClient(apiKey string, apiEndpoint string) (*Client, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("missing apiKey")
	}
	if len(apiEndpoint) == 0 {
		return nil, errors.New("missing apiEndpoint")
	}

	endpoint, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, err
	}

	c := &Client{apiKey: apiKey, endpoint: endpoint}
	if c.HttpClient == nil {
		c.HttpClient = cleanhttp.DefaultClient()
	}

	return c, nil
}

const packageName = "github.com/fnaoto/go_deploygate"

var userAgent = fmt.Sprintf("GoDeployGate (+%s; %s)", packageName, runtime.Version())

func (c *Client) Get(httpRequest *HttpRequest) (*http.Response, error) {
	httpRequest.method = "GET"
	return c.NewRequest(httpRequest)
}

func (c *Client) Post(httpRequest *HttpRequest) (*http.Response, error) {
	httpRequest.method = "POST"
	return c.NewRequest(httpRequest)
}

func (c *Client) Patch(httpRequest *HttpRequest) (*http.Response, error) {
	httpRequest.method = "PATCH"
	return c.NewRequest(httpRequest)
}

func (c *Client) Delete(httpRequest *HttpRequest) (*http.Response, error) {
	httpRequest.method = "DELETE"
	return c.NewRequest(httpRequest)
}

func (c *Client) NewRequest(httpRequest *HttpRequest) (*http.Response, error) {
	u := *c.endpoint
	u.Path = path.Join(c.endpoint.Path, httpRequest.path)

	req, err := http.NewRequest(httpRequest.method, u.String(), httpRequest.body)
	if err != nil {
		return nil, err
	}

	if httpRequest.header == nil {
		httpRequest.header = &Header{
			accept:      "application/json",
			contentType: "application/x-www-form-urlencoded",
		}
	}

	req.Header.Set("Accept", httpRequest.header.accept)
	req.Header.Set("Content-Type", httpRequest.header.contentType)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.apiKey))

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Decode(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode >= 300 {
		return fmt.Errorf("status code is %v, body is %v", resp.StatusCode, buf.String())
	}

	if resp.ContentLength == 0 {
		return nil
	}

	err := json.Unmarshal(buf.Bytes(), &out)

	if err != nil {
		return fmt.Errorf("%v, output type is %T, body is %v", err, out, buf.String())
	}

	return nil
}
