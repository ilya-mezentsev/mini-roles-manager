package permission

import (
	"github.com/jmoiron/sqlx"
	"mini-roles-backend/source/domains/permission/spec"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

const (
	createRecursivePermissionsFunctionQuery = `
	create or replace function recursive_permissions(
		roles_version_id character(32),
		entry_point_role_id character(32),
		_account_hash character(32),
		depth int,
		exclude character(32)[]
	)
	returns table(permissions character(32)[], permissions_depth int)
	language plpgsql
	as $$
	declare
	    extended_role_id character(32);
	begin
		return query select
			r.permissions permissions,
			depth permissions_depth
		from role r where role_id = entry_point_role_id and account_hash = _account_hash and version_id = roles_version_id;

		for extended_role_id in
			(
			    select unnest(extends) from role
			    where
			          role_id = entry_point_role_id and
			          account_hash = _account_hash and
			          version_id = roles_version_id
			)
		loop
		    if not extended_role_id = any(exclude) then
				return query select * from recursive_permissions(
				    roles_version_id,
					trim(extended_role_id),
					_account_hash,
					depth + 1,
					array_append(exclude, extended_role_id)
				) rp;
		    end if;
		end loop;
	end $$;`

	selectPermissionsQuery = `
	select
	       trim(p.permission_id) permission_id,
	       trim(p.operation) operation,
	       trim(p.effect) effect,
	       trim(r.resource_id) resource_id,
	       r.links_to links_to
	from permission p
	inner join resource r on r.resource_id = p.resource_id
	cross join lateral (select * from recursive_permissions($2, $3, $1, 1, array[]::character(32)[])) rp
	where permission_id = any(rp.permissions)
	order by rp.permissions_depth`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) List(spec spec.PermissionWithAccountIdAndRoleId) ([]sharedModels.Permission, error) {
	_, err := r.db.Exec(createRecursivePermissionsFunctionQuery)
	if err != nil {
		return nil, err
	}

	var proxies []permissionProxy
	err = r.db.Select(&proxies, selectPermissionsQuery, spec.AccountId, spec.RolesVersionId, spec.RoleId)

	return r.makePermissions(proxies), err
}

func (r Repository) makePermissions(proxies []permissionProxy) []sharedModels.Permission {
	var permissions []sharedModels.Permission
	for _, proxy := range proxies {
		permissions = append(permissions, sharedModels.Permission{
			Id: proxy.PermissionId,
			Resource: sharedModels.Resource{
				Id:      proxy.ResourceId,
				LinksTo: proxy.makeLinksTo(),
			},
			Operation: proxy.Operation,
			Effect:    proxy.Effect,
		})
	}

	return permissions
}
