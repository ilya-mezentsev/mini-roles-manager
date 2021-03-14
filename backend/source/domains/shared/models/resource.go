package models

type (
	Resource struct {
		Id      ResourceId   `json:"resource_id" db:"resource_id"`
		Title   string       `json:"title" db:"title"`
		LinksTo []ResourceId `json:"links_to" db:"links_to"`
	}
)
