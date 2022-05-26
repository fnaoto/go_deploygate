package deploygate

import (
	"bytes"
	"fmt"
)

func (c *Client) ListOrganizationTeamMembers(req *ListOrganizationTeamMembersRequest) (*ListOrganizationTeamMembersResponse, error) {
	path := fmt.Sprintf("organizations/%s/teams/%s/users", req.Organization, req.Team)

	resp, err := c.Get(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *ListOrganizationTeamMembersResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) AddOrganizationTeamMember(req *AddOrganizationTeamMemberRequest) (*AddOrganizationTeamMemberResponse, error) {
	path := fmt.Sprintf("organizations/%s/teams/%s/users", req.Organization, req.Team)
	body := fmt.Sprintf("user=%s", req.User)

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *AddOrganizationTeamMemberResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
