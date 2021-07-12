package opensearch

// TenantAction wraps possible values for tenant permissions
// via constant values
type TenantAction string

const (
	// RoleAlertingAckAlerts Grants permissions to view and acknowledge alerts, but not modify destinations or monitors.
	RoleAlertingAckAlerts string = "alerting_ack_alerts"
	// RoleAlertingFullAccess Grants full permissions to all alerting actions.
	RoleAlertingFullAccess string = "alerting_full_access"
	// RoleAlertingReadAccess Grants permissions to view alerts, destinations, and monitors, but not acknowledge alerts or modify destinations or monitors.
	RoleAlertingReadAccess string = "alerting_read_access"
	// RoleAnomalyFullAccess Grants full permissions to all anomaly detection actions.
	RoleAnomalyFullAccess string = "anomaly_full_access"
	// RoleAnomalyReadAccess Grants permissions to view detectors, but not create, modify, or delete detectors.
	RoleAnomalyReadAccess string = "anomaly_read_access"
	// RoleAllAccess Grants full access to the cluster: all cluster-wide operations, write to all indices, write to all tenants.
	RoleAllAccess string = "all_access"
	// RoleKibanaReadOnly A special role that prevents users from making changes to visualizations, dashboards, and other OpenSearch Dashboards objects.
	RoleKibanaReadOnly string = "kibana_read_only"
	// RoleKibanaUser Grants permissions to use OpenSearch Dashboards: cluster-wide searches, index monitoring, and write to various OpenSearch Dashboards indices.
	RoleKibanaUser string = "kibana_user"
	// RoleLogstash Grants permissions for Logstash to interact with the cluster: cluster-wide searches, cluster monitoring, and write to the various Logstash indices.
	RoleLogstash string = "logstash"
	// RoleManageSnapshots Grants permissions to manage snapshot repositories, take snapshots, and restore snapshots.
	RoleManageSnapshots string = "manage_snapshots"
	// RoleReadall Grants permissions for cluster-wide searches like msearch and search permissions for all indices.
	RoleReadall string = "readall"
	// RoleReadallAndMonitor Same as readall, but with added cluster monitoring permissions.
	RoleReadallAndMonitor string = "readall_and_monitor"
	// RoleSecurityRestApiAccess A special role that allows access to the REST API. See plugins.security.restapi.roles_enabled in opensearch.yml and Access control for the API.
	RoleSecurityRestApiAccess string = "security_rest_api_access"
	// RoleReportsReadAccess Grants permissions to generate on-demand reports, download existing reports, and view report definitions, but not to create report definitions.
	RoleReportsReadAccess string = "reports_read_access"
	// RoleReportsInstancesReadAccess Grants permissions to generate on-demand reports and download existing reports, but not to view or create report definitions.
	RoleReportsInstancesReadAccess string = "reports_instances_read_access"
	// RoleReportsFullAccess Grants full permissions to reports.
	RoleReportsFullAccess string = "reports_full_access"
	// RoleAsynchronousSearchFullAccess Grants full permissions to all asynchronous search actions.
	RoleAsynchronousSearchFullAccess string = "asynchronous_search_full_access"
	// RoleAsynchronousSearchReadAccess Grants permissions to view asynchronous searches, but not to submit, modify, or delete async searches.
	RoleAsynchronousSearchReadAccess string = "asynchronous_search_read_access"

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
