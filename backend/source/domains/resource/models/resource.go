package models

import shared "mini-roles-backend/source/domains/shared/models"

type (
	Resource struct {
		Id      shared.ResourceId   `json:"resource_id" db:"resource_id"`
		Title   string              `json:"title" db:"title"`
		LinksTo []shared.ResourceId `json:"links_to" db:"links_to"`
	}
)
