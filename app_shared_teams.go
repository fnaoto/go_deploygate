package deploygate

import "fmt"

func (c *Client) ListAppSharedTeams(cfg *ListAppSharedTeamsRequest) (*ListAppSharedTeamsResponse, error) {
	path := fmt.Sprintf("organizations/%s/platforms/%s/apps/%s/shared_teams", cfg.Organizations, cfg.Platform, cfg.AppId)

	resp, err := c.Get(&HttpRequest{
		spath: path,
	})
	if err != nil {
		return nil, err
	}

	var r *ListAppSharedTeamsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
