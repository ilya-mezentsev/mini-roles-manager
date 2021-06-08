package shared

import (
	"fmt"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/shared/services/hash"
)

func MakePassword(c models.AccountCredentials) string {
	return hash.Md5(fmt.Sprintf("%s:%s", c.Password, c.Password))
}
