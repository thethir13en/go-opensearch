package opensearch

// TenantAction wraps possible values for tenant permissions
// via constant values
type TenantAction string

const (
	// ActionGroups
	ActionGroupRead                        string = "read"
	ActionGroupClusterAll                  string = "cluster_all"
	ActionGroupUnlimited                   string = "unlimited"
	ActionGroupCreateIndex                 string = "create_index"
	ActionGroupIndex                       string = "index"
	ActionGroupClusterCompositeOpsRo       string = "cluster_composite_ops_ro"
	ActionGroupSuggest                     string = "suggest"
	ActionGroupDataAccess                  string = "data_access"
	ActionGroupDelete                      string = "delete"
	ActionGroupClusterManagePipelines      string = "cluster_manage_pipelines"
	ActionGroupManageAliases               string = "manage_aliases"
	ActionGroupManage                      string = "manage"
	ActionGroupCrud                        string = "crud"
	ActionGroupManageSnapshots             string = "manage_snapshots"
	ActionGroupKibanaAllRead               string = "kibana_all_read"
	ActionGroupSearch                      string = "search"
	ActionGroupClusterCompositeOps         string = "cluster_composite_ops"
	ActionGroupGet                         string = "get"
	ActionGroupClusterMonitor              string = "cluster_monitor"
	ActionGroupWrite                       string = "write"
	ActionGroupIndicesMonitor              string = "indices_monitor"
	ActionGroupClusterManageIndexTemplates string = "cluster_manage_index_templates"
	ActionGroupIndicesAll                  string = "indices_all"
	ActionGroupKibanaAllWrite              string = "kibana_all_write"

	// Tenant Actions
	TenantActionKibanaAllRead  TenantAction = "kibana_all_read"
	TenantActionKibanaAllWrite TenantAction = "kibana_all_write"
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
	TenantPatterns []string       `json:"tenant_patterns"`
	AllowedActions []TenantAction `json:"allowed_actions,omitempty"`
}
