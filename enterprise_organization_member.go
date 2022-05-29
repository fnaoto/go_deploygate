package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) ListEnterpriseOrganizationMembers(req *ListEnterpriseOrganizationMembersRequest) (*ListEnterpriseOrganizationMembersResponse, error) {
	path := fmt.Sprintf("enterprises/%s/organizations/%s/users", req.Enterprise, req.Organization)

	resp, err := c.Get(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *ListEnterpriseOrganizationMembersResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) AddEnterpriseOrganizationMembers(req *AddEnterpriseOrganizationMembersRequest) (*AddEnterpriseOrganizationMembersResponse, error) {
	path := fmt.Sprintf("enterprises/%s/organizations/%s/users", req.Enterprise, req.Organization)
	body := fmt.Sprintf("user=%s", req.User)

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *AddEnterpriseOrganizationMembersResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) RemoveEnterpriseOrganizationMembers(req *RemoveEnterpriseOrganizationMembersRequest) (*RemoveEnterpriseOrganizationMembersResponse, error) {
	path := fmt.Sprintf("enterprises/%s/organizations/%s/users/%s", req.Enterprise, req.Organization, req.User)

	resp, err := c.Delete(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *RemoveEnterpriseOrganizationMembersResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
