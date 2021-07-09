package opensearch

// TenantPermissionAction wraps possible values for tenant permissions
// via constant values
type TenantPermissionAction string

const (
	TenantPermissionActionKibanaAllRead  TenantPermissionAction = "kibana_all_read"
	TenantPermissionActionKibanaAllWrite TenantPermissionAction = "kibana_all_write"
)

// RoleDefinition provides stucture which represents json of a role definition
//
type RoleDefinition struct {
	Reserved           bool               `json:"reserved,omitempty"`
	Hidden             bool               `json:"hidden,omitempty"`
	Description        string             `json:"description,omitempty"`
	ClusterPermissions []string           `json:"cluster_permissions,omitempty"`
	IndexPermissions   []IndexPermission  `json:"index_permissions,omitempty"`
	TenantPermissions  []TenantPermission `json:"tenant_permissions,omitempty"`
	Static             bool               `json:"static,omitempty"`
}

// IndexPermission provides stucture which represents json of an index permission definition
//
type IndexPermission struct {
	IndexPatterns  []string `json:"index_patterns"`
	Fls            []string `json:"fls,omitempty"`
	Dls            string   `json:"dls,omitempty"`
	MaskedFields   []string `json:"masked_fields,omitempty"`
	AllowedActions []string `json:"allowed_actions,omitempty"`
}

// TenantPermission provides stucture which represents json of a tenant permission definition
//
type TenantPermission struct {
	TenantPatterns []string                 `json:"tenant_patterns"`
	AllowedActions []TenantPermissionAction `json:"allowed_actions,omitempty"`
}