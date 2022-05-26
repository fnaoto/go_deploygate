package deploygate

import "fmt"

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
