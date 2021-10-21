package permission

const CreatePermissionQuery = `
insert into permission(resource_id, account_hash, permission_id, operation, effect)
values(:resource_id, :account_hash, :permission_id, :operation, :effect)`
