package deploygate

import (
	"fmt"
)

func (c *Client) GetAppMembers(cfg *AppMemberConfig) (*GetAppMemberResponse, error) {
	path := fmt.Sprintf("/users/%s/platforms/%s/apps/%s/members", cfg.Owner, cfg.Platform, cfg.AppId)

	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var r *GetAppMemberResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
