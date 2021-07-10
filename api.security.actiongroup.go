package opensearch

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
)
