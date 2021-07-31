package roles_version

import (
	"github.com/jmoiron/sqlx"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

const (
	selectRolesVersionQuery = `
	select trim(version_id) version_id, trim(title) title
	from roles_version where account_hash = $1 order by created_at`

	addRolesVersionQuery = `insert into roles_version(version_id, title, account_hash) values(:version_id, :title, :account_hash)`

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
	_, err := r.db.NamedExec(addRolesVersionQuery, r.mapFromRolesVersion(accountId, rolesVersion))
	if sharedError.IsDuplicateKey(err) {
		err = sharedError.DuplicateUniqueKey{}
	}

	return err
}

func (r Repository) mapFromRolesVersion(
	accountId sharedModels.AccountId,
	rolesVersion sharedModels.RolesVersion,
) map[string]interface{} {
	return map[string]interface{}{
		"account_hash": accountId,
		"version_id":   rolesVersion.Id,
		"title":        rolesVersion.Title,
	}
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
	_, err := r.db.NamedExec(updateRolesVersionQuery, r.mapFromRolesVersion(accountId, rolesVersion))

	return err
}

func (r Repository) Delete(
	accountId sharedModels.AccountId,
	rolesVersionId sharedModels.RolesVersionId,
) error {
	_, err := r.db.Exec(deleteRolesVersionQuery, accountId, rolesVersionId)

	return err
}
