package session_check

import (
	responseFactory "github.com/ilya-mezentsev/response-factory"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/domains/account/interfaces"
	"mini-roles-backend/source/domains/account/request"
	"mini-roles-backend/source/domains/account/services/shared"
	"mini-roles-backend/source/domains/account/spec"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	sharedKeys "mini-roles-backend/source/shared/keys"
)

type Service struct {
	repository interfaces.SessionRepository
}

func New(
	repository interfaces.SessionRepository,
) Service {
	return Service{
		repository: repository,
	}
}

// CheckSessionFromCookie used for middleware. nil return result is meaning that request can be processed
func (s Service) CheckSessionFromCookie(request request.SessionExists) responseFactory.Response {
	cookieToken, err := request.Context.Cookie(shared.CookieTokenKey)
	if err != nil {
		return responseFactory.UnauthorizedError(sharedError.ServiceError{
			Code:        missedTokenInCookiesCode,
			Description: missedTokenInCookiesDescription,
		})
	}

	return s.checkSession(request, cookieToken)
}

func (s Service) checkSession(request request.SessionExists, token string) responseFactory.Response {
	accountExists, err := s.repository.SessionExists(spec.SessionWithId{
		Id: sharedModels.AccountId(token),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"account_id": token,
		}).Errorf("Unable to check account existance: %v", err)
		return response_factory.DefaultServerError()
	}

	if accountExists {
		request.Context.Set(sharedKeys.ContextTokenKey, token)
		return nil
	} else {
		return responseFactory.ForbiddenError(sharedError.ServiceError{
			Code:        noAccountByTokenCode,
			Description: noAccountByTokenDescription,
		})
	}
}

// CheckSessionFromHeader used for middleware. nil return result is meaning that request can be processed
func (s Service) CheckSessionFromHeader(request request.SessionExists) responseFactory.Response {
	tokenFromHeader := request.Context.GetHeader(headerTokenKey)
	if tokenFromHeader == "" {
		return responseFactory.UnauthorizedError(sharedError.ServiceError{
			Code:        missedTokenInHeadersCode,
			Description: missedTokenInHeadersDescription,
		})
	}

	return s.checkSession(request, tokenFromHeader)
}
