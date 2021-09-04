package registration

import (
	"errors"
	"fmt"
	responseFactory "github.com/ilya-mezentsev/response-factory"
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
	registrationRepository        interfaces.RegistrationRepository
	rolesVersionCreatorRepository sharedInterfaces.RolesVersionCreatorRepository
}

func New(
	registrationRepository interfaces.RegistrationRepository,
	rolesVersionCreatorRepository sharedInterfaces.RolesVersionCreatorRepository,
) Service {
	return Service{
		registrationRepository:        registrationRepository,
		rolesVersionCreatorRepository: rolesVersionCreatorRepository,
	}
}

func (s Service) Register(request request.Registration) responseFactory.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	request.Credentials.Password = shared.MakePassword(request.Credentials)
	session := s.createSession(request)
	err := s.registrationRepository.Register(
		session,
		request.Credentials,
	)
	if err != nil {
		if errors.As(err, &sharedError.DuplicateUniqueKey{}) {
			return responseFactory.ClientError(sharedError.ServiceError{
				Code:        shared.LoginAlreadyExistsCode,
				Description: shared.LoginAlreadyExistsDescription,
			})
		}

		log.Errorf("Unable to register user in DB: %v", err)
		return response_factory.DefaultServerError()
	}

	err = s.rolesVersionCreatorRepository.Create(
		session.Id,
		sharedModels.RolesVersion{
			Id: defaultRolesVersionId,
		},
	)
	if err != nil {
		log.WithFields(log.Fields{
			"account_hash": session.Id,
		}).Errorf("Unable to create roles version: %v", err)
		return response_factory.DefaultServerError()
	}

	return responseFactory.DefaultResponse()
}

func (s Service) createSession(request request.Registration) models.AccountSession {
	return models.AccountSession{
		Id: sharedModels.AccountId(
			hash.Md5WithTimeAsKey(fmt.Sprintf("%s:%s", request.Credentials.Login, request.Credentials.Password)),
		),
	}
}
