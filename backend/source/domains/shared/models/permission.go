package models

type (
	Permission struct {
		Id        PermissionId `json:"permission_id" db:"permission_id"`
		Resource  Resource     `json:"resource_id"`
		Operation string       `json:"operation" db:"operation"` // create|read|update|delete
		Effect    string       `json:"effect" db:"effect"`       // permit|deny
	}
)
