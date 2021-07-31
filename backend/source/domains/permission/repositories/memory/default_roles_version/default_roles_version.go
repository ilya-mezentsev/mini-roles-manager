package default_roles_version

import (
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

type Repository struct {
	appData sharedModels.AppData
}

func New(appData sharedModels.AppData) Repository {
	return Repository{appData}
}

func (r Repository) Fetch(sharedSpec.AccountWithId) (sharedModels.RolesVersion, error) {
	return sharedModels.RolesVersion{
		Id: r.appData.DefaultRolesVersionId,
	}, nil
}
