package presenter

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/shared/interfaces"
	"net/http"
)

func MakeJsonResponse(c *gin.Context, r interfaces.Response) {
	c.Status(r.HttpStatus())
	if r.HasData() {
		c.JSON(r.HttpStatus(), gin.H{
			"status": r.ApplicationStatus(),
			"data":   r.GetData(),
		})
	}
}

func MakeInvalidJsonResponse(c *gin.Context) {
	c.String(http.StatusBadRequest, "Invalid JSON format")
}
