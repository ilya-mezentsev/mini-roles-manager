package resource

var (
	resourcesOperations        = [...]string{"create", "read", "update", "delete"}
	resourcesOperationsEffects = [...]string{"permit", "deny"}
)

const (
	resourceExistsCode        = "resource-already-exists"
	resourceExistsDescription = "Resource with provided id is already exist"
)
