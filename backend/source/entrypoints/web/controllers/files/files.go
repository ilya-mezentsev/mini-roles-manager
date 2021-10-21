package files

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/files/request"
	"mini-roles-backend/source/domains/files/services/export"
	"mini-roles-backend/source/domains/files/services/import_file"
	"mini-roles-backend/source/entrypoints/web/shared/context_keys"
	"mini-roles-backend/source/entrypoints/web/shared/presenter"
)

type Controller struct {
	exportService export.Service
	importService import_file.Service
}

func New(
	exportService export.Service,
	importService import_file.Service,
) Controller {
	return Controller{
		exportService: exportService,
		importService: importService,
	}
}

func (c Controller) Export(context *gin.Context) {
	r := request.ExportRequest{
		AccountId: context_keys.GetAccountId(context),
	}

	presenter.MakeFileResponse(
		context,
		c.exportService.MakeExportFile(r),
	)
}

func (c Controller) Import(context *gin.Context) {
	file, _, err := context.Request.FormFile("app_data_file")
	if err != nil {
		presenter.MakeBadRequestResponse(context)
		return
	}

	presenter.MakeJsonResponse(
		context,
		c.importService.ImportFromFile(request.ImportRequest{
			AccountId: context_keys.GetAccountId(context),
			File:      file,
		}),
	)
}
