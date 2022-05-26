package deploygate

import "fmt"

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
