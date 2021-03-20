package header

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/account/request"
	"mini-roles-backend/source/domains/account/services/session_check"
	"mini-roles-backend/source/entrypoints/web/shared/presenter"
)

type Middleware struct {
	service session_check.Service
}

func New(service session_check.Service) Middleware {
	return Middleware{service}
}

func (m Middleware) HasSessionInHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		interceptResponse := m.service.CheckSessionFromHeader(request.SessionExists{
			Context: c,
		})
		if interceptResponse != nil {
			presenter.MakeJsonResponse(c, interceptResponse)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
