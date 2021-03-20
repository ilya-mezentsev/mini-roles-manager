package session

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/domains/account/interfaces"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/response_factory"
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

func (s Service) CreateSession(request request.CreateSession) sharedInterfaces.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	accountSession, err := s.repository.GetSession(request.Credentials)
	if err != nil {
		if errors.As(err, &sharedError.EntryNotFound{}) {
			return response_factory.ClientError(sharedError.ServiceError{
				Code:        credentialsNotFoundCode,
				Description: credentialsNotFoundDescription,
			})
		}

		log.Errorf("Unable to fetch session from DB: %v", err)
		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	s.setCookie(request.Context, accountSession.Id)
	return response_factory.SuccessResponse(accountSession)
}

func (s Service) setCookie(c *gin.Context, value sharedModels.AccountId) {
	c.SetCookie(
		cookieTokenKey,
		string(value),
		cookieMaxAge,
		cookiePath,
		s.config.ServerDomain(),
		s.config.ServerSecureCookie(),
		cookieHttpOnly,
	)
}

func (s Service) GetSession(request request.GetSession) sharedInterfaces.Response {
	cookieToken, err := request.Context.Cookie(cookieTokenKey)
	if err != nil {
		return response_factory.DefaultResponse()
	}

	accountSession := models.AccountSession{
		Id: sharedModels.AccountId(cookieToken),
	}
	accountExists, err := s.repository.SessionExists(accountSession)
	if err != nil {
		log.Errorf("Unable to check account existance: %v", err)
		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	if accountExists {
		s.setCookie(request.Context, accountSession.Id)
		return response_factory.SuccessResponse(accountSession)
	} else {
		return response_factory.DefaultResponse()
	}
}

func (s Service) DeleteSession(request request.DeleteSession) sharedInterfaces.Response {
	s.unsetCookie(request.Context)
	return response_factory.DefaultResponse()
}

func (s Service) unsetCookie(c *gin.Context) {
	c.SetCookie(
		cookieTokenKey,
		cookieUnsetValue,
		cookieUnsetMaxAge,
		cookiePath,
		s.config.ServerDomain(),
		s.config.ServerSecureCookie(),
		cookieHttpOnly,
	)
}
