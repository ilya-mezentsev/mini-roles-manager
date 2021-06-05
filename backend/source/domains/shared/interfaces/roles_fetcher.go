package interfaces

import shared "mini-roles-backend/source/domains/shared/models"

type RolesFetcherRepository interface {
	List(accountId shared.AccountId) ([]shared.Role, error)
}
