package mock

import (
	shared "mini-roles-backend/source/domains/shared/models"
	sharedResource "mini-roles-backend/source/domains/shared/resource"
)

const (
	PermittedOperation        = sharedResource.ReadOperation
	DeniedOperation           = sharedResource.DeleteOperation
	DefinedOnLinkingOperation = sharedResource.UpdateOperation
	UndefinedOperation        = sharedResource.CreateOperation
	LinkingResourceId         = shared.ResourceId("linking-resource-id")
)
