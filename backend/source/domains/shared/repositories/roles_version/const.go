package roles_version

const (
	AddRolesVersionQuery = `insert into roles_version(version_id, title, account_hash) values(:version_id, :title, :account_hash)`
)
