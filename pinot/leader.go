package pinot

import (
	"fmt"
)

var leaderEndpoint = "leader"

// GetLeadersForTables returns the leaders of the tables
func (c *Client) GetLeadersForTables() (*map[string]interface{}, error) {
	leaderURL := fmt.Sprintf("%s/tables", leaderEndpoint)
	var res map[string]interface{}
	err := c.get(leaderURL, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetLeaderForTable returns the leader of the table
func (c *Client) GetLeaderForTable(name string) (*map[string]interface{}, error) {
	leaderURL := fmt.Sprintf("%s/tables/%s", leaderEndpoint, name)
	var res map[string]interface{}
	err := c.get(leaderURL, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
