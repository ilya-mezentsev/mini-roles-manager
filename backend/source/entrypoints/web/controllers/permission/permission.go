package permission

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/permission/request"
	"mini-roles-backend/source/domains/permission/services/permission"
	"mini-roles-backend/source/entrypoints/web/shared/context_keys"
	"mini-roles-backend/source/entrypoints/web/shared/presenter"
)

type Controller struct {
	service permission.Service
}

func New(service permission.Service) Controller {
	return Controller{service}
}

func (c Controller) ResolveResourceAccessEffect(context *gin.Context) {
	var r request.PermissionAccess
	if err := context.ShouldBindJSON(&r); err != nil {
		presenter.MakeInvalidJsonResponse(context)
		return
	}

	presenter.MakeJsonResponse(
		context,
		c.service.HasPermission(
			context_keys.GetAccountId(context),
			r,
		),
	)
}
