package interfaces

import (
	shared "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

type RolesFetcherRepository interface {
	List(spec sharedSpec.AccountWithId) ([]shared.Role, error)
}
