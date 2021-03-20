package error

type DuplicateUniqueKey struct {
}

func (d DuplicateUniqueKey) Error() string {
	return "DuplicateUniqueKey"
}

type EntryNotFound struct {
}

func (d EntryNotFound) Error() string {
	return "EntryNotFound"
}
