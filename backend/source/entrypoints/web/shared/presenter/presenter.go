package presenter

import (
	"github.com/gin-gonic/gin"
	responseFactory "github.com/ilya-mezentsev/response-factory"
	"net/http"
)

func MakeJsonResponse(c *gin.Context, r responseFactory.Response) {
	c.Status(r.HttpStatus())
	if r.HasData() {
		c.JSON(r.HttpStatus(), gin.H{
			"status": applicationStatus(r),
			"data":   r.Data(),
		})
	}
}

func applicationStatus(r responseFactory.Response) string {
	if r.IsOk() {
		return statusOk
	} else {
		return statusError
	}
}

func MakeInvalidJsonResponse(c *gin.Context) {
	c.String(http.StatusBadRequest, "Invalid JSON format")
}

func MakeBadRequestResponse(c *gin.Context) {
	c.String(http.StatusBadRequest, "Bad request")
}

func MakeInternalErrorResponse(c *gin.Context) {
	c.String(http.StatusInternalServerError, "Internal server error")
}

func MakeFileResponse(c *gin.Context, r responseFactory.Response) {
	c.Status(r.HttpStatus())
	if r.HasData() {
		c.File(r.Data().(string))
	}
}
