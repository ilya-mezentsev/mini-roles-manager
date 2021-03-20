package context_keys

import (
	"github.com/gin-gonic/gin"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedKeys "mini-roles-backend/source/shared/keys"
)

func GetAccountId(context *gin.Context) sharedModels.AccountId {
	return sharedModels.AccountId(context.GetString(sharedKeys.ContextTokenKey))
}
