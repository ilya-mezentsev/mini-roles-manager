package models

type (
	Resource struct {
		Id      ResourceId   `json:"resource_id"`
		Title   string       `json:"title"`
		LinksTo []ResourceId `json:"links_to"`
	}
)
