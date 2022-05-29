package deploygate

import (
	"bytes"
	"fmt"
	"net/url"
)

func (c *Client) ListEnterpriseSharedTeams(req *ListEnterpriseSharedTeamsRequest) (*ListEnterpriseSharedTeamsResponse, error) {
	path := fmt.Sprintf("enterprises/%s/shared_teams", req.Enterprise)

	resp, err := c.Get(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *ListEnterpriseSharedTeamsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) AddEnterpriseSharedTeam(req *AddEnterpriseSharedTeamRequest) (*AddEnterpriseSharedTeamResponse, error) {
	path := fmt.Sprintf("enterprises/%s/shared_teams", req.Enterprise)
	body := fmt.Sprintf("name=%s&description=%s", req.SharedTeam, url.QueryEscape(req.Description))

	resp, err := c.Post(&HttpRequest{
		path: path,
		body: bytes.NewBufferString(body),
	})
	if err != nil {
		return nil, err
	}

	var r *AddEnterpriseSharedTeamResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) RemoveEnterpriseSharedTeam(req *RemoveEnterpriseSharedTeamRequest) (*RemoveEnterpriseSharedTeamResponse, error) {
	path := fmt.Sprintf("enterprises/%s/shared_teams/%s", req.Enterprise, req.SharedTeam)

	resp, err := c.Delete(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *RemoveEnterpriseSharedTeamResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
