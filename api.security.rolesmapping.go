package opensearch

// RolesMappingDefinition provides stucture which represents json of a role mapping definition
//
type RolesMappingDefinition struct {
	Reserved        bool     `json:"reserved,omitempty"`
	Hidden          bool     `json:"hidden,omitempty"`
	Description     string   `json:"description,omitempty"`
	Hosts           []string `json:"hosts,omitempty"`
	Users           []string `json:"users,omitempty"`
	BackendRoles    []string `json:"backend_roles,omitempty"`
	AndBackendRoles []string `json:"and_backend_roles,omitempty"`
}
