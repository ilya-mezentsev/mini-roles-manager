package mock

import shared "mini-roles-backend/source/domains/shared/models"

const (
	PermittedOperation        = "read"
	DeniedOperation           = "delete"
	DefinedOnLinkingOperation = "update"
	UndefinedOperation        = "create"
	LinkingResourceId         = shared.ResourceId("linking-resource-id")
)
