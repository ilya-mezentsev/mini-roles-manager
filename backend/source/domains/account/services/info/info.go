package info

import (
	"errors"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/domains/account/interfaces"
	"mini-roles-backend/source/domains/account/models"
	"mini-roles-backend/source/domains/account/request"
	"mini-roles-backend/source/domains/account/services/shared"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	"mini-roles-backend/source/domains/shared/services/response_factory"
)

type Service struct {
	repository interfaces.AccountInfoRepository
}

func New(repository interfaces.AccountInfoRepository) Service {
	return Service{repository}
}

func (s Service) GetInfo(request request.GetInfoRequest) sharedInterfaces.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	info, err := s.repository.FetchInfo(request.AccountId)
	if err != nil {
		log.Errorf("Unable to fetch account info: %v", err)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	return response_factory.SuccessResponse(info)
}

func (s Service) UpdateCredentials(request request.UpdateCredentialsRequest) sharedInterfaces.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	if request.Credentials.Password == "" {
		err = s.repository.UpdateLogin(request.AccountId, request.Credentials.Login)
	} else {
		request.Credentials.Password = shared.MakePassword(models.AccountCredentials(request.Credentials))
		err = s.repository.UpdateCredentials(request.AccountId, request.Credentials)
	}

	if err != nil {
		if errors.As(err, &sharedError.DuplicateUniqueKey{}) {
			return response_factory.ClientError(sharedError.ServiceError{
				Code:        shared.LoginAlreadyExistsCode,
				Description: shared.LoginAlreadyExistsDescription,
			})
		}

		log.Errorf("Unable to update account password: %v", err)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	return response_factory.DefaultResponse()
}
