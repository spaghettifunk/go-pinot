package pinot

// Version returns the version number of Pinot components
func (c *Client) Version() (*map[string]string, error) {
	var res map[string]string
	err := c.get("version", &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
