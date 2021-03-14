package models

type (
	Permission struct {
		Id        PermissionId `json:"permission_id"`
		Resource  Resource     `json:"resource_id"`
		Operation string       `json:"operation"`          // create|read|update|delete
		Effect    string       `json:"effect" db:"effect"` // permit|deny
	}
)
