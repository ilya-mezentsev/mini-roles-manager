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
			"status": r.ApplicationStatus(),
			"data":   r.Data(),
		})
	}
}

func MakeInvalidJsonResponse(c *gin.Context) {
	c.String(http.StatusBadRequest, "Invalid JSON format")
}

func MakeFileResponse(c *gin.Context, r responseFactory.Response) {
	c.Status(r.HttpStatus())
	if r.HasData() {
		c.File(r.Data().(string))
	}
}
