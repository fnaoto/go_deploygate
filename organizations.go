package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) ListOrganizations() (*ListOrganizationsResponse, error) {
	path := "/organizations"

	resp, err := c.Get(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *ListOrganizationsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) CreateOrganizations(req *CreateOrganizationsRequest) (*CreateOrganizationsResponse, error) {
	path := "/organizations"
	body := fmt.Sprintf("name=%s&description=%s", req.Name, req.Description)

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *CreateOrganizationsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
