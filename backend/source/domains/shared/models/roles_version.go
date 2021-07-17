package models

type (
	RolesVersion struct {
		Id    RolesVersionId `json:"id" db:"version_id" validate:"required"`
		Title string         `json:"title" db:"title"`
	}
)
