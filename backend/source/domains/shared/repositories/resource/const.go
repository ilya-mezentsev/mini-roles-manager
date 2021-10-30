package resource

const (
	CreateResourceQuery = `
	insert into resource(account_hash, resource_id, title, links_to)
	values(:account_hash, :resource_id, :title, :links_to)`
)
