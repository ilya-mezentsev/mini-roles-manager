package registration

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/domains/account/interfaces"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/request"
	"mini-roles-backend/source/domains/account/services/shared"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/hash"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"mini-roles-backend/source/domains/shared/services/validation"
)

type Service struct {
	repository interfaces.RegistrationRepository
}

func New(repository interfaces.RegistrationRepository) Service {
	return Service{repository}
}

func (s Service) Register(request request.Registration) sharedInterfaces.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	request.Credentials.Password = shared.MakePassword(request.Credentials)
	err := s.repository.Register(
		s.createSession(request),
		request.Credentials,
	)
	if err != nil {
		if errors.As(err, &sharedError.DuplicateUniqueKey{}) {
			return response_factory.ClientError(sharedError.ServiceError{
				Code:        shared.LoginAlreadyExistsCode,
				Description: shared.LoginAlreadyExistsDescription,
			})
		}

		log.Errorf("Unable to register user in DB: %v", err)
		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	return response_factory.DefaultResponse()
}

func (s Service) createSession(request request.Registration) models.AccountSession {
	return models.AccountSession{
		Id: sharedModels.AccountId(
			hash.Md5WithTimeAsKey(fmt.Sprintf("%s:%s", request.Credentials.Login, request.Credentials.Password)),
		),
	}
}
