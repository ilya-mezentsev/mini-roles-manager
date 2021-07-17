package permission

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/permission/request"
	"mini-roles-backend/source/domains/permission/services/permission"
	sharedModels "mini-roles-backend/source/domains/shared/models"
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
	requestQuery := context.Request.URL.Query()
	r := request.PermissionAccess{
		RoleId:         sharedModels.RoleId(requestQuery.Get("roleId")),
		ResourceId:     sharedModels.ResourceId(requestQuery.Get("resourceId")),
		Operation:      requestQuery.Get("operation"),
		RolesVersionId: sharedModels.RolesVersionId(requestQuery.Get("rolesVersionId")),
	}

	presenter.MakeJsonResponse(
		context,
		c.service.HasPermission(
			context_keys.GetAccountId(context),
			r,
		),
	)
}
