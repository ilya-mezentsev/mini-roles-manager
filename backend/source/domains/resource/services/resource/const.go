package resource

import sharedResource "mini-roles-backend/source/domains/shared/resource"

var (
	resourcesOperations = [...]string{
		sharedResource.CreateOperation,
		sharedResource.ReadOperation,
		sharedResource.UpdateOperation,
		sharedResource.DeleteOperation,
	}
	resourcesOperationsEffects = [...]string{
		sharedResource.PermitEffect,
		sharedResource.DenyEffect,
	}
)

const (
	resourceExistsCode        = "resource-already-exists"
	resourceExistsDescription = "Resource with provided id is already exist"
)
