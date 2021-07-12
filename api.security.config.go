package opensearch

// FilteredAliasMode wraps possible values for filtered_alias_mode
// via constant values
type FilteredAliasMode string

const (
	FilteredAliasModeDisallow FilteredAliasMode = "disallow"
	FilteredAliasModeWarn     FilteredAliasMode = "warn"
	FilteredAliasModeNoWarn   FilteredAliasMode = "nowarn"
)

// SecurityConfig defines security plugin configuration
//
type SecurityConfig struct {
	Dynamic SecurityConfigDynamic `json:"dynamic"`
}

// SecurityConfigDynamic defines dynamic field of security
// plugin configuration
//
type SecurityConfigDynamic struct {
	FilteredAliasMode    FilteredAliasMode `json:"filtered_alias_mode,omitempty"`
	DoNotFailOnForbidden bool              `json:"do_not_fail_on_forbidden,omitempty"`

	Kibana SecurityConfigKibana           `json:"kibana,omitempty"`
	HTTP   SecurityConfigHTTP             `json:"http,omitempty"`
	Authc  map[string]SecurityConfigAuthc `json:"authc,omitempty"`
	Authz  map[string]SecurityConfigAuthz `json:"authz,omitempty"`
}

// SecurityConfigKibana defines kibana parameters of security configuration
//
type SecurityConfigKibana struct {
	MultitenancyEnabled bool   `json:"multitenancy_enabled,omitempty"`
	ServerUsername      string `json:"server_username,omitempty"`
	Index               string `json:"index,omitempty"`
}

// SecurityConfigHTTP defines http parameters of security configuraion
//
type SecurityConfigHTTP struct {
	AnonymusAuthEnabled bool `json:"anonymus_auth_enabled,omitempty"`

	Xff SecurityConfigHTTPXff `json:"xff,omitempty"`
}

// SecurityConfigHTTPXff defines http/xff parameters of security configuraion
//
type SecurityConfigHTTPXff struct {
	Enabled         bool   `json:"enabled,omitempty"`
	InternalProxies string `json:"internalProxies,omitempty"`
	RemoteIpHeader  string `json:"remoteIpHeader,omitempty"`
}

// SecurityConfigAuthc defines authentication backend parameters of security configuration
//
type SecurityConfigAuthc struct {
	Description      string `json:"description,omitempty"`
	HttpEnabled      bool   `json:"http_enabled,omitempty"`
	TransportEnabled bool   `json:"transport_enabled,omitempty"`
	Order            int    `json:"order"`

	HttpAuthenticator     SecurityConfigHTTPAuthenticator     `json:"http_authenticator,omitempty"`
	AuthenticationBackend SecurityConfigBackend `json:"authentication_backend,omitempty"`
}

// SecurityConfigHTTPAuthenticator defines authc/http_authenticatior parameters
//
type SecurityConfigHTTPAuthenticator struct {
	Type      string                 `json:"type"`
	Config    map[string]interface{} `json:"config,omitempty"`
	Challenge bool                   `json:"challenge,omitempty"`
}

// SecurityConfigAuthz defines authorization backend parameters of security configuration
//
type SecurityConfigAuthz struct {
	Description      string `json:"description,omitempty"`
	HttpEnabled      bool   `json:"http_enabled,omitempty"`
	TransportEnabled bool   `json:"transport_enabled,omitempty"`

	AuthorizationBackend SecurityConfigBackend `json:"authorization_backend,omitempty"`
}

// SecurityConfigBackend defines backend configuration for authc or authz
//
type SecurityConfigBackend struct {
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config,omitempty"`
}
