package request

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/account/models"
)

type (
	CreateSession struct {
		Context     *gin.Context
		Credentials models.AccountCredentials `json:"credentials" validate:"required"`
	}

	GetSession struct {
		Context *gin.Context
	}

	DeleteSession struct {
		Context *gin.Context
	}

	SessionExists struct {
		Context *gin.Context
	}
)
