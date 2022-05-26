package deploygate

import (
	"bytes"
	"fmt"
	"net/url"
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

func (c *Client) CreateOrganization(req *CreateOrganizationRequest) (*CreateOrganizationResponse, error) {
	path := "/organizations"
	body := fmt.Sprintf("name=%s&description=%s", req.Name, url.QueryEscape(req.Description))

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *CreateOrganizationResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) GetOrganization(req *GetOrganizationRequest) (*GetOrganizationResponse, error) {
	path := fmt.Sprintf("/organizations/%s", req.Name)

	resp, err := c.Get(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *GetOrganizationResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) UpdateOrganization(req *UpdateOrganizationRequest) (*UpdateOrganizationResponse, error) {
	path := fmt.Sprintf("/organizations/%s", req.Name)
	body := fmt.Sprintf("description=%s", url.QueryEscape(req.Description))

	resp, err := c.Patch(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *UpdateOrganizationResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) DeleteOrganization(req *DeleteOrganizationRequest) (*DeleteOrganizationResponse, error) {
	path := fmt.Sprintf("/organizations/%s", req.Name)

	resp, err := c.Delete(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *DeleteOrganizationResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
