package error

import "strings"

func IsDuplicateKey(err error) bool {
	return err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}
