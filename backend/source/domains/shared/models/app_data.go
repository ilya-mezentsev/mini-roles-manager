package models

type AppData struct {
	Resources []Resource `json:"resources"`
	Roles     []Role     `json:"roles"`
}
