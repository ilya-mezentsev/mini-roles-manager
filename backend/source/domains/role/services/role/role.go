package role

import (
	"errors"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/domains/role/interfaces"
	"mini-roles-backend/source/domains/role/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	shared "mini-roles-backend/source/domains/shared/interfaces"
	"mini-roles-backend/source/domains/shared/services/response_factory"
)

type Service struct {
	repository interfaces.RoleRepository
}

func New(repository interfaces.RoleRepository) Service {
	return Service{repository}
}

func (s Service) Create(request request.CreateRoleRequest) shared.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	err = s.repository.Create(request.AccountId, request.Role)
	if err != nil {
		if errors.As(err, &sharedError.DuplicateUniqueKey{}) {
			return response_factory.ClientError(sharedError.ServiceError{
				Code:        roleExistsCode,
				Description: roleExistsDescription,
			})
		}

		log.WithFields(log.Fields{
			"role": request.Role,
		}).Errorf("Unable to create role: %v", err)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	return response_factory.DefaultResponse()
}

func (s Service) RolesList(request request.RolesListRequest) shared.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	roles, err := s.repository.List(request.AccountId)
	if err != nil {
		log.Errorf("Unable to fetch roles from DB: %v", err)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	return response_factory.SuccessResponse(roles)
}

func (s Service) UpdateRole(request request.UpdateRoleRequest) shared.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	err = s.repository.Update(request.AccountId, request.Role)
	if err != nil {
		log.WithFields(log.Fields{
			"role": request.Role,
		}).Errorf("Unable to update role: %v", err)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	return response_factory.DefaultResponse()
}

func (s Service) DeleteRole(request request.DeleteRoleRequest) shared.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	err = s.repository.Delete(request.AccountId, request.RoleId)
	if err != nil {
		log.Errorf("Unable to delete role: %v", err)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	return response_factory.DefaultResponse()
}
