package account

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/request"
	"mini-roles-backend/source/domains/account/services/registration"
	"mini-roles-backend/source/domains/account/services/session"
	"mini-roles-backend/source/entrypoints/web/shared/presenter"
)

type Controller struct {
	registrationService registration.Service
	sessionService      session.Service
}

func New(
	registrationService registration.Service,
	sessionService session.Service,
) Controller {
	return Controller{
		registrationService: registrationService,
		sessionService:      sessionService,
	}
}

func (c Controller) Register(context *gin.Context) {
	var r request.Registration
	if err := context.ShouldBindJSON(&r); err != nil {
		presenter.MakeInvalidJsonResponse(context)
		return
	}

	presenter.MakeJsonResponse(
		context,
		c.registrationService.Register(r),
	)
}

func (c Controller) Login(context *gin.Context) {
	presenter.MakeJsonResponse(
		context,
		c.sessionService.GetSession(request.GetSession{
			Context: context,
		}),
	)
}

func (c Controller) SignIn(context *gin.Context) {
	var credentials models.AccountCredentials
	if err := context.ShouldBindJSON(&credentials); err != nil {
		presenter.MakeInvalidJsonResponse(context)
		return
	}

	presenter.MakeJsonResponse(
		context,
		c.sessionService.CreateSession(request.CreateSession{
			Context:     context,
			Credentials: credentials,
		}),
	)
}

func (c Controller) SignOut(context *gin.Context) {
	presenter.MakeJsonResponse(
		context,
		c.sessionService.DeleteSession(request.DeleteSession{
			Context: context,
		}),
	)
}
