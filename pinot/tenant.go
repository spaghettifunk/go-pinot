package pinot

import (
	"fmt"

	"github.com/spaghettifunk/go-pinot/pkg/models"
	"github.com/spaghettifunk/go-pinot/pkg/util"
)

var tenantsEndpoint = "/tenants"

// GetTenant retrieve the specified tenant
func (c *Client) GetTenant(tenant *models.Tenant) (*models.TenantInstances, error) {
	tenantURL := fmt.Sprintf("%s/%s?type=%s", tenantsEndpoint, util.PointerToString(tenant.Name), tenant.Role)
	var tr models.TenantInstances
	err := c.get(tenantURL, &tr)
	if err != nil {
		return nil, err
	}
	return &tr, nil
}

// CreateTenant creates a tenant
func (c *Client) CreateTenant(tenant *models.Tenant) (*models.TenantInstances, error) {
	tenantURL := fmt.Sprintf("%s", tenantsEndpoint)
	var tr models.TenantInstances
	err := c.post(tenantURL, tenant, &tr)
	if err != nil {
		return nil, err
	}
	return &tr, nil
}

// UpdateTenant updates the tenant entity
func (c *Client) UpdateTenant(tenant *models.Tenant) (*models.TenantInstances, error) {
	tenantURL := fmt.Sprintf("%s/%s?type=%s", tenantsEndpoint, util.PointerToString(tenant.Name), tenant.Role)
	var tr models.TenantInstances
	err := c.put(tenantURL, tenant, &tr)
	if err != nil {
		return nil, err
	}
	return &tr, nil
}

// DeleteTenant deletes the tenant entity
func (c *Client) DeleteTenant(tenant *models.Tenant) (*models.TenantInstances, error) {
	tenantURL := fmt.Sprintf("%s/%s?type=%s", tenantsEndpoint, util.PointerToString(tenant.Name), tenant.Role)
	var tr models.TenantInstances
	err := c.delete(tenantURL, &tr)
	if err != nil {
		return nil, err
	}
	return &tr, nil
}

// GetAllTenants retrieves all the tenants available
func (c *Client) GetAllTenants() (*models.TenantInstances, error) {
	tenantURL := fmt.Sprintf("%s", tenantsEndpoint)
	var tr models.TenantInstances
	err := c.get(tenantURL, &tr)
	if err != nil {
		return nil, err
	}
	return &tr, nil
}

// GetTenantTables retrieve the specified tenant's tables
func (c *Client) GetTenantTables(tenant *models.Tenant) (*models.TenantTables, error) {
	tenantURL := fmt.Sprintf("%s/%s/tables", tenantsEndpoint, util.PointerToString(tenant.Name))
	var tt models.TenantTables
	err := c.get(tenantURL, &tt)
	if err != nil {
		return nil, err
	}
	return &tt, nil
}

// ChangeTenantState changes the state of the tenant
func (c *Client) ChangeTenantState(tenant *models.Tenant, state models.State) (*map[string]map[string]interface{}, error) {
	tenantURL := fmt.Sprintf("%s/%s/tables?type=%s&state=%s", tenantsEndpoint, util.PointerToString(tenant.Name), tenant.Role, state)
	var tr map[string]map[string]interface{}
	err := c.post(tenantURL, tenant, &tr)
	if err != nil {
		return nil, err
	}
	return &tr, nil
}

// GetTenantMetadata retrieves the specified tenant's metadata
func (c *Client) GetTenantMetadata(tenant *models.Tenant) (*models.TenantInstances, error) {
	tenantURL := fmt.Sprintf("%s/%s/metadata?type=%s", tenantsEndpoint, util.PointerToString(tenant.Name), tenant.Role)
	var tr models.TenantInstances
	err := c.get(tenantURL, &tr)
	if err != nil {
		return nil, err
	}
	return &tr, nil
}
