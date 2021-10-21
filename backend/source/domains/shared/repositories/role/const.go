package role

const (
	AddRoleQuery = `
	insert into role(account_hash, role_id, version_id, title, permissions, extends)
	values(:account_hash, :role_id, :version_id, :title, :permissions, :extends)`
)
