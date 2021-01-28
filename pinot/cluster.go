package pinot

import (
	"encoding/json"
	"fmt"
)

var clusterEndpoint = "/cluster"

// ClusterInfo holds the name of the cluster
type ClusterInfo struct {
	Name string `json:"clusterName"`
}

// GetClusterConfigurations returns the available configurations
func (c *Client) GetClusterConfigurations() (*map[string]interface{}, error) {
	clusterURL := fmt.Sprintf("%s/configs", clusterEndpoint)
	var cf map[string]interface{}
	err := c.get(clusterURL, &cf)
	if err != nil {
		return nil, err
	}
	return &cf, nil
}

// UpdateClusterConfiguration updates the configurations of the cluster
func (c *Client) UpdateClusterConfiguration(configs json.RawMessage) (*string, error) {
	clusterURL := fmt.Sprintf("%s/configs", clusterEndpoint)
	var res string
	err := c.post(clusterURL, configs, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteClusterConfiguration deletes a config
func (c *Client) DeleteClusterConfiguration(property string) (*string, error) {
	clusterURL := fmt.Sprintf("%s/configs/%s", clusterEndpoint, property)
	var res string
	err := c.delete(clusterURL, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetClusterInfo returns the info of the cluster
func (c *Client) GetClusterInfo() (*ClusterInfo, error) {
	clusterURL := fmt.Sprintf("%s/info", clusterEndpoint)
	var ci ClusterInfo
	err := c.get(clusterURL, &ci)
	if err != nil {
		return nil, err
	}
	return &ci, nil
}
