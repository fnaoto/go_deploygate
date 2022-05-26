package deploygate

func (c *Client) ListOrganizations() (*ListOrganizationsResponse, error) {
	path := "/organizations"

	resp, err := c.Get(&HttpRequest{
		path: path,
	})
	if err != nil {
		return nil, err
	}

	var r *ListOrganizationsResponse
	err = c.Decode(resp, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
