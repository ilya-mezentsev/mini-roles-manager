package session

import (
	"errors"
	"github.com/gin-gonic/gin"
	responseFactory "github.com/ilya-mezentsev/response-factory"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/domains/account/interfaces"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/request"
	"mini-roles-backend/source/domains/account/services/shared"
	"mini-roles-backend/source/domains/account/spec"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"mini-roles-backend/source/domains/shared/services/validation"
)

type Service struct {
	repository interfaces.SessionRepository
	config     config.ServerConfigsRepository
}

func New(
	repository interfaces.SessionRepository,
	config config.ServerConfigsRepository,
) Service {
	return Service{
		repository: repository,
		config:     config,
	}
}

func (s Service) CreateSession(request request.CreateSession) responseFactory.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	request.Credentials.Password = shared.MakePassword(request.Credentials)
	accountSession, err := s.repository.GetSession(spec.SessionWithCredentials{
		Credentials: request.Credentials,
	})
	if err != nil {
		if errors.As(err, &sharedError.EntryNotFound{}) {
			return responseFactory.ClientError(sharedError.ServiceError{
				Code:        credentialsNotFoundCode,
				Description: credentialsNotFoundDescription,
			})
		}

		log.Errorf("Unable to fetch session from DB: %v", err)
		return response_factory.DefaultServerError()
	}

	s.setCookie(request.Context, accountSession.Id)
	return responseFactory.SuccessResponse(accountSession)
}

func (s Service) setCookie(c *gin.Context, value sharedModels.AccountId) {
	c.SetCookie(
		shared.CookieTokenKey,
		string(value),
		cookieMaxAge,
		cookiePath,
		s.config.ServerDomain(),
		s.config.ServerSecureCookie(),
		cookieHttpOnly,
	)
}

func (s Service) GetSession(request request.GetSession) responseFactory.Response {
	cookieToken, err := request.Context.Cookie(shared.CookieTokenKey)
	if err != nil {
		return responseFactory.DefaultResponse()
	}

	accountSession := models.AccountSession{
		Id: sharedModels.AccountId(cookieToken),
	}
	accountExists, err := s.repository.SessionExists(spec.SessionWithId(accountSession))
	if err != nil {
		log.WithFields(log.Fields{
			"account_id": accountSession.Id,
		}).Errorf("Unable to check account existance: %v", err)
		return response_factory.DefaultServerError()
	}

	if accountExists {
		s.setCookie(request.Context, accountSession.Id)
		return responseFactory.SuccessResponse(accountSession)
	} else {
		return responseFactory.DefaultResponse()
	}
}

func (s Service) DeleteSession(request request.DeleteSession) responseFactory.Response {
	s.unsetCookie(request.Context)
	return responseFactory.DefaultResponse()
}

func (s Service) unsetCookie(c *gin.Context) {
	c.SetCookie(
		shared.CookieTokenKey,
		cookieUnsetValue,
		cookieUnsetMaxAge,
		cookiePath,
		s.config.ServerDomain(),
		s.config.ServerSecureCookie(),
		cookieHttpOnly,
	)
}
