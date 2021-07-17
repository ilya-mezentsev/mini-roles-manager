package role

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/role/request"
	"mini-roles-backend/source/domains/role/services/role"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/entrypoints/web/shared/context_keys"
	"mini-roles-backend/source/entrypoints/web/shared/presenter"
)

type Controller struct {
	service role.Service
}

func New(service role.Service) Controller {
	return Controller{service}
}

func (c Controller) CreateRole(context *gin.Context) {
	var r request.CreateRole
	if err := context.ShouldBindJSON(&r); err != nil {
		presenter.MakeInvalidJsonResponse(context)
		return
	}

	r.AccountId = context_keys.GetAccountId(context)
	presenter.MakeJsonResponse(
		context,
		c.service.Create(r),
	)
}

func (c Controller) RolesList(context *gin.Context) {
	r := request.RolesList{
		AccountId: context_keys.GetAccountId(context),
	}

	presenter.MakeJsonResponse(
		context,
		c.service.RolesList(r),
	)
}

func (c Controller) UpdateRole(context *gin.Context) {
	var r request.UpdateRole
	if err := context.ShouldBindJSON(&r); err != nil {
		presenter.MakeInvalidJsonResponse(context)
		return
	}

	r.AccountId = context_keys.GetAccountId(context)
	presenter.MakeJsonResponse(
		context,
		c.service.UpdateRole(r),
	)
}

func (c Controller) DeleteRole(context *gin.Context) {
	r := request.DeleteRole{
		AccountId:      context_keys.GetAccountId(context),
		RoleId:         sharedModels.RoleId(context.Param("role_id")),
		RolesVersionId: sharedModels.RolesVersionId(context.Param("roles_version_id")),
	}

	presenter.MakeJsonResponse(
		context,
		c.service.DeleteRole(r),
	)
}
