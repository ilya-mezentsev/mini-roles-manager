package models

type AppData struct {
	Resources             []Resource     `json:"resources"`
	Roles                 []Role         `json:"roles"`
	DefaultRolesVersionId RolesVersionId `json:"default_roles_version_id"`
}
