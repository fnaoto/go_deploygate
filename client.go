package deploygate

import (
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
	c, err := NewClient(os.Getenv(DGApiTokenEnv))
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewClient(apiKey string) (*Client, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("missing apiKey")
	}
	c := &Client{apiKey: apiKey}
	return c.init()
}

func (c *Client) init() (*Client, error) {
	e, err := url.Parse(DGApiEndpoint)
	if err != nil {
		return nil, err
	}
	c.endpoint = e

	if c.httpClient == nil {
		c.httpClient = cleanhttp.DefaultClient()
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

func (c *Client) Delete(httpRequest *HttpRequest) (*http.Response, error) {
	httpRequest.method = "DELETE"
	return c.NewRequest(httpRequest)
}

func (c *Client) NewRequest(httpRequest *HttpRequest) (*http.Response, error) {
	u := *c.endpoint
	u.Path = path.Join(c.endpoint.Path, httpRequest.path)

	q := u.Query()
	q.Set("token", c.apiKey)
	u.RawQuery = q.Encode()

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

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Decode(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("status code is %v", resp.StatusCode)
	}
	if resp.ContentLength == 0 {
		return nil
	}
	decodeer := json.NewDecoder(resp.Body)
	err := decodeer.Decode(out)
	if err != nil {
		return fmt.Errorf("cloudn't decode json: %v", resp.Body)
	}
	return nil
}
