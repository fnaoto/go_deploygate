package deploygate

import "fmt"

func (c *Client) DeleteDistributionsPage(req *DeleteDistributionsPageRequest) (*DeleteDistributionsPageResponse, error) {
	path := fmt.Sprintf("/distributions/%s", req.Distribution)

	resp, err := c.Delete(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *DeleteDistributionsPageResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
