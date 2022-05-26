package deploygate

import (
	"bytes"
	"fmt"
	"net/url"
)

func (c *Client) ListOrganizationMembers(req *ListOrganizationMembersRequest) (*ListOrganizationMembersResponse, error) {
	path := fmt.Sprintf("/organizations/%s/members", req.Organization)

	resp, err := c.Get(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *ListOrganizationMembersResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) AddOrganizationMemberByUserName(req *AddOrganizationMemberByUserNameRequest) (*AddOrganizationMemberResponse, error) {
	path := fmt.Sprintf("/organizations/%s/members", req.Organization)
	body := fmt.Sprintf("username=%s", req.UserName)

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *AddOrganizationMemberResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) AddOrganizationMemberByEmail(req *AddOrganizationMemberByEmailRequest) (*AddOrganizationMemberResponse, error) {
	path := fmt.Sprintf("/organizations/%s/members", req.Organization)
	body := fmt.Sprintf("email=%s", url.QueryEscape(req.Email))

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *AddOrganizationMemberResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
