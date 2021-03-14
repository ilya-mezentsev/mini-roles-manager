package error

type DuplicateUniqueKey struct {
}

func (d DuplicateUniqueKey) Error() string {
	return "DuplicateUniqueKey"
}
