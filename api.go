package opensearch

import "github.com/elastic/go-elasticsearch/v7/esapi"

// API provides access to Opendistro APIs
//
type API struct {
	Security *Security
}

// NewAPI returns new instance of Opensearch API
//
func NewAPI(t esapi.Transport) *API {
	return &API{
		Security: &Security{
			DeleteActionGroup:  newDeleteActionGroupFunc(t),
			DeleteInternalUser: newDeleteInternalUserFunc(t),
			DeleteRole:         newDeleteRoleFunc(t),
			DeleteRolesMapping: newDeleteRolesMappingFunc(t),
			GetActionGroup:     newGetActionGroupFunc(t),
			GetInternalUser:    newGetInternalUserFunc(t),
			GetRole:            newGetRoleFunc(t),
			GetRolesMapping:    newGetRolesMappingFunc(t),
			PatchRole:          newPatchRoleFunc(t),
			PostRole:           newPostRoleFunc(t),
			PutRole:            newPutRoleFunc(t),
		},
	}
}

// Security allows access to Opendistro Security plugin APIs
//
type Security struct {
	DeleteActionGroup  DeleteActionGroup
	DeleteInternalUser DeleteInternalUser
	DeleteRole         DeleteRole
	DeleteRolesMapping DeleteRolesMapping
	GetActionGroup     GetActionGroup
	GetInternalUser    GetInternalUser
	GetRole            GetRole
	GetRolesMapping    GetRolesMapping
	PatchRole          PatchRole
	PostRole           PostRole
	PutRole            PutRole
}
