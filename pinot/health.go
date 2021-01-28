package pinot

// Health checks if the cluster is alive
func (c *Client) Health() (*string, error) {
	var res string
	err := c.get("health", &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// ControllerHealth checks the controller health
func (c *Client) ControllerHealth() (*string, error) {
	var res string
	err := c.get("pinot-controller/admin", &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
