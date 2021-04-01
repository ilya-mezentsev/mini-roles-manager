package models

type (
	Permission struct {
		Id        PermissionId `json:"id"`
		Operation string       `json:"operation"` // create|read|update|delete
		Effect    string       `json:"effect"`    // permit|deny
		Resource  Resource     `json:"-"`
	}
)
