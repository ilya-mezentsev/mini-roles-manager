package mock

import (
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedResource "mini-roles-backend/source/domains/shared/resource"
)

const (
	PermittedOperation        = sharedResource.ReadOperation
	DeniedOperation           = sharedResource.DeleteOperation
	DefinedOnLinkingOperation = sharedResource.UpdateOperation
	UndefinedOperation        = sharedResource.CreateOperation
	LinkingResourceId         = sharedModels.ResourceId("linking-resource-id")
)
