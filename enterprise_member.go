package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) ListEnterpriseMembers(req *ListEnterpriseMembersRequest) (*ListEnterpriseMembersResponse, error) {
	path := fmt.Sprintf("/enterprises/%s/users", req.Enterprise)

	resp, err := c.Get(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *ListEnterpriseMembersResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) AddEnterpriseMember(req *AddEnterpriseMemberRequest) (*AddEnterpriseMemberResponse, error) {
	path := fmt.Sprintf("/enterprises/%s/users", req.Enterprise)
	body := fmt.Sprintf("user=%s", req.User)

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *AddEnterpriseMemberResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) RemoveEnterpriseMember(req *RemoveEnterpriseMemberRequest) (*RemoveEnterpriseMemberResponse, error) {
	path := fmt.Sprintf("/enterprises/%s/users/%s", req.Enterprise, req.User)

	resp, err := c.Delete(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *RemoveEnterpriseMemberResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
