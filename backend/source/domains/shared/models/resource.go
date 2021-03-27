package models

type (
	Resource struct {
		Id          ResourceId   `json:"id" validate:"required"`
		Title       string       `json:"title"`
		LinksTo     []ResourceId `json:"linksTo"`
		Permissions []Permission `json:"permissions"`
	}
)
