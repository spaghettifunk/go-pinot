package models

// Role is a custom type to describe the tenant role
type Role string

// State is a custom type to describe the state of the tenant
type State string

const (
	// BROKER is the broker role
	BROKER Role = "BROKER"
	// SERVER is the server role
	SERVER Role = "SERVER"
	// ENABLE state
	ENABLE State = "enable"
	// DISABLE state
	DISABLE State = "disable"
	// DROP state
	DROP State = "drop"
)

// Tenant represents the definition of the tenant model in apache pinot
type Tenant struct {
	Role              Role    `json:"tenantRole" validate:"required,role"`
	Name              *string `json:"tenantName" validate:"required"`
	NumberOfInstances *int32  `json:"numberOfInstances,omitempty"`
	OfflineInstances  *int32  `json:"offileInstances,omitempty"`
	RealtimeInstances *int32  `json:"realtimeInstances,omitempty"`
}

// TenantInstances holds the broker/server instances
type TenantInstances struct {
	Name            string   `json:"tenantName"`
	BrokerInstances []string `json:"BrokerInstances,omitempty"`
	ServerInstances []string `json:"ServerInstances,omitempty"`
}

// TenantTables holds the tables belonging to a tenant
type TenantTables struct {
	Tables []string `json:"tables"`
}
