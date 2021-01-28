package pinot

import (
	"fmt"

	"github.com/spaghettifunk/pinot-go-client/models"
)

var schemasEndpoint = "schemas"

// GetSchemas returns all the schemas' name of the cluster
func (c *Client) GetSchemas() (*[]string, error) {
	schemasURL := fmt.Sprintf("%s", schemasEndpoint)
	var ss []string
	err := c.get(schemasURL, &ss)
	if err != nil {
		return nil, err
	}
	return &ss, nil
}

// GetSchema returns the schema for a given name
func (c *Client) GetSchema(name string) (*models.Schema, error) {
	schemaURL := fmt.Sprintf("%s/%s", schemasEndpoint, name)
	var ss models.Schema
	err := c.get(schemaURL, &ss)
	if err != nil {
		return nil, err
	}
	return &ss, nil
}

// CreateSchema creates a new schema in the cluster
func (c *Client) CreateSchema(schema *models.Schema, override bool) (*map[string]string, error) {
	schemasURL := fmt.Sprintf("%s?override=%t", schemasEndpoint, override)
	var res map[string]string
	err := c.post(schemasURL, schema, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteSchema deletes a schema in the cluster
func (c *Client) DeleteSchema(name string) (*map[string]string, error) {
	schemasURL := fmt.Sprintf("%s/%s", schemasEndpoint, name)
	var res map[string]string
	err := c.delete(schemasURL, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetTableSchema returns the schema of a selected table
func (c *Client) GetTableSchema(tableName string) (*models.Schema, error) {
	schemasURL := fmt.Sprintf("tables/%s/schema", tableName)
	var ss models.Schema
	err := c.get(schemasURL, &ss)
	if err != nil {
		return nil, err
	}
	return &ss, nil
}

// ValidateSchema validates the schema
func (c *Client) ValidateSchema(fdmp models.FormDataMultiPart) (*map[string]interface{}, error) {
	schemasURL := fmt.Sprintf("%s/validate", schemasEndpoint)
	var res map[string]interface{}
	err := c.post(schemasURL, fdmp, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateSchema updates the current schema
func (c *Client) UpdateSchema(name string, fdmp models.FormDataMultiPart) (*map[string]interface{}, error) {
	schemasURL := fmt.Sprintf("%s/%s", schemasEndpoint, name)
	var res map[string]interface{}
	err := c.put(schemasURL, fdmp, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
