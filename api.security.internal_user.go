package opensearch

// InternalUserDefinition provides stucture which represents json of a user definition
//
type InternalUserDefinition struct {
	Reserved                bool              `json:"reserved,omitempty"`
	Hidden                  bool              `json:"hidden,omitempty"`
	Hash                    string            `json:"hash,omitempty"`
	Password                string            `json:"password,omitempty"`
	Description             string            `json:"description,omitempty"`
	BackendRoles            []string          `json:"backend_roles,omitempty"`
	OpendistroSecurityRoles []string          `json:"opendistro_secrutiy_roles,omitempty"`
	Attributes              map[string]string `json:"attributes,omitempty"`
	Static                  bool              `json:"static,omitempty"`
}
