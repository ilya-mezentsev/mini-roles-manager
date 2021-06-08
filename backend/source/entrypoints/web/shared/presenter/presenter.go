package presenter

import (
	"github.com/gin-gonic/gin"
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	"net/http"
)

func MakeJsonResponse(c *gin.Context, r sharedInterfaces.Response) {
	c.Status(r.HttpStatus())
	if r.HasData() {
		c.JSON(r.HttpStatus(), gin.H{
			"status": r.ApplicationStatus(),
			"data":   r.Data(),
		})
	}
}

func MakeInvalidJsonResponse(c *gin.Context) {
	c.String(http.StatusBadRequest, "Invalid JSON format")
}

func MakeFileResponse(c *gin.Context, r sharedInterfaces.Response) {
	c.Status(r.HttpStatus())
	if r.HasData() {
		c.File(r.Data().(string))
	}
}
