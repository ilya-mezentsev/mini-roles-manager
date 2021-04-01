package permission

import (
	"github.com/jmoiron/sqlx"
	sharedModels "mini-roles-backend/source/domains/shared/models"
)

const (
	createRecursivePermissionsFunctionQuery = `
	create or replace function recursive_permissions(
		entry_point_role_id character(32),
		_account_hash character(32),
		depth int
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
		from role r where role_id = entry_point_role_id and account_hash = _account_hash;

		for extended_role_id in
			(select unnest(extends) from role where role_id = entry_point_role_id and account_hash = _account_hash)
		loop
			return query select * from recursive_permissions(
			    trim(extended_role_id),
			    _account_hash,
			    depth + 1
			) rp;
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
	cross join lateral (select * from recursive_permissions($2, $1, 1)) rp
	where permission_id = any(rp.permissions)
	order by rp.permissions_depth`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db}
}

func (r Repository) List(
	accountId sharedModels.AccountId,
	roleId sharedModels.RoleId,
) ([]sharedModels.Permission, error) {
	_, err := r.db.Exec(createRecursivePermissionsFunctionQuery)
	if err != nil {
		return nil, err
	}

	var proxies []permissionProxy
	err = r.db.Select(&proxies, selectPermissionsQuery, accountId, roleId)

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
