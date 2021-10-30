package roles_version

import (
	"github.com/jmoiron/sqlx"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedRolesVersionRepository "mini-roles-backend/source/domains/shared/repositories/roles_version"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

const (
	selectRolesVersionQuery = `
	select trim(version_id) version_id, trim(title) title
	from roles_version where account_hash = $1 order by created_at`

	updateRolesVersionQuery = `
	update roles_version set title = :title
	where account_hash = :account_hash and version_id = :version_id`

	deleteRolesVersionQuery = `delete from roles_version where account_hash = $1 and version_id = $2`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) Create(
	accountId sharedModels.AccountId,
	rolesVersion sharedModels.RolesVersion,
) error {
	_, err := r.db.NamedExec(
		sharedRolesVersionRepository.AddRolesVersionQuery,
		sharedRolesVersionRepository.MapFromRolesVersion(accountId, rolesVersion),
	)
	if sharedError.IsDuplicateKey(err) {
		err = sharedError.DuplicateUniqueKey{}
	}

	return err
}

func (r Repository) List(spec sharedSpec.AccountWithId) ([]sharedModels.RolesVersion, error) {
	var rolesVersions []sharedModels.RolesVersion
	err := r.db.Select(&rolesVersions, selectRolesVersionQuery, spec.AccountId)

	return rolesVersions, err
}

func (r Repository) Update(
	accountId sharedModels.AccountId,
	rolesVersion sharedModels.RolesVersion,
) error {
	_, err := r.db.NamedExec(
		updateRolesVersionQuery,
		sharedRolesVersionRepository.MapFromRolesVersion(accountId, rolesVersion),
	)

	return err
}

func (r Repository) Delete(
	accountId sharedModels.AccountId,
	rolesVersionId sharedModels.RolesVersionId,
) error {
	_, err := r.db.Exec(deleteRolesVersionQuery, accountId, rolesVersionId)

	return err
}
