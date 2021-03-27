package resource

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/resource/request"
	"mini-roles-backend/source/domains/resource/services/resource"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/entrypoints/web/shared/context_keys"
	"mini-roles-backend/source/entrypoints/web/shared/presenter"
)

type Controller struct {
	service resource.Service
}

func New(service resource.Service) Controller {
	return Controller{service}
}

func (c Controller) CreateResource(context *gin.Context) {
	var r request.CreateResource
	if err := context.ShouldBindJSON(&r); err != nil {
		presenter.MakeInvalidJsonResponse(context)
		return
	}

	r.AccountId = context_keys.GetAccountId(context)
	presenter.MakeJsonResponse(
		context,
		c.service.CreateResource(r),
	)
}

func (c Controller) ResourcesList(context *gin.Context) {
	r := request.ResourcesList{
		AccountId: context_keys.GetAccountId(context),
	}

	presenter.MakeJsonResponse(
		context,
		c.service.ResourcesList(r),
	)
}

func (c Controller) UpdateResource(context *gin.Context) {
	var r request.UpdateResource
	if err := context.ShouldBindJSON(&r); err != nil {
		presenter.MakeInvalidJsonResponse(context)
		return
	}

	r.AccountId = context_keys.GetAccountId(context)
	presenter.MakeJsonResponse(
		context,
		c.service.UpdateResource(r),
	)
}

func (c Controller) DeleteResource(context *gin.Context) {
	r := request.DeleteResource{
		AccountId:  context_keys.GetAccountId(context),
		ResourceId: sharedModels.ResourceId(context.Param("resource_id")),
	}

	presenter.MakeJsonResponse(
		context,
		c.service.DeleteResource(r),
	)
}
