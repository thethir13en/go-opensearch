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
			GetConfig:          newGetConfigFunc(t),
			GetInternalUser:    newGetInternalUserFunc(t),
			GetRole:            newGetRoleFunc(t),
			GetRolesMapping:    newGetRolesMappingFunc(t),
			PatchRole:          newPatchRoleFunc(t),
			PatchRolesMapping:  newPatchRolesMappingFunc(t),
			PostRole:           newPostRoleFunc(t),
			PostRolesMapping:   newPostRolesMappingFunc(t),
			PutRole:            newPutRoleFunc(t),
			PutRolesMapping:    newPutRolesMappingFunc(t),
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
	GetConfig          GetConfig
	GetInternalUser    GetInternalUser
	GetRole            GetRole
	GetRolesMapping    GetRolesMapping
	PatchRole          PatchRole
	PatchRolesMapping  PatchRolesMapping
	PostRole           PostRole
	PostRolesMapping   PostRolesMapping
	PutRole            PutRole
	PutRolesMapping    PutRolesMapping
}
