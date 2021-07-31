package roles_version

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/role/request"
	"mini-roles-backend/source/domains/role/services/roles_version"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/entrypoints/web/shared/context_keys"
	"mini-roles-backend/source/entrypoints/web/shared/presenter"
)

type Controller struct {
	service roles_version.Service
}

func New(service roles_version.Service) Controller {
	return Controller{service}
}

func (c Controller) CreateRolesVersion(context *gin.Context) {
	var r request.CreateRolesVersion
	if err := context.ShouldBindJSON(&r); err != nil {
		presenter.MakeInvalidJsonResponse(context)
		return
	}

	r.AccountId = context_keys.GetAccountId(context)
	presenter.MakeJsonResponse(
		context,
		c.service.CreateRolesVersion(r),
	)
}

func (c Controller) RolesVersionsList(context *gin.Context) {
	r := request.RolesVersionList{
		AccountId: context_keys.GetAccountId(context),
	}

	presenter.MakeJsonResponse(
		context,
		c.service.RolesVersionList(r),
	)
}

func (c Controller) UpdateRolesVersion(context *gin.Context) {
	var r request.UpdateRolesVersion
	if err := context.ShouldBindJSON(&r); err != nil {
		presenter.MakeInvalidJsonResponse(context)
		return
	}

	r.AccountId = context_keys.GetAccountId(context)
	presenter.MakeJsonResponse(
		context,
		c.service.UpdateRolesVersion(r),
	)
}

func (c Controller) DeleteRolesVersion(context *gin.Context) {
	r := request.DeleteRolesVersion{
		AccountId:      context_keys.GetAccountId(context),
		RolesVersionId: sharedModels.RolesVersionId(context.Param("roles_version_id")),
	}

	presenter.MakeJsonResponse(
		context,
		c.service.DeleteRolesVersion(r),
	)
}
