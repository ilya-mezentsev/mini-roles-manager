package files

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/files/request"
	"mini-roles-backend/source/domains/files/services/export"
	"mini-roles-backend/source/entrypoints/web/shared/context_keys"
	"mini-roles-backend/source/entrypoints/web/shared/presenter"
)

type Controller struct {
	service export.Service
}

func New(service export.Service) Controller {
	return Controller{service}
}

func (c Controller) Export(context *gin.Context) {
	r := request.ExportRequest{
		AccountId: context_keys.GetAccountId(context),
	}

	presenter.MakeFileResponse(
		context,
		c.service.MakeExportFile(r),
	)
}
