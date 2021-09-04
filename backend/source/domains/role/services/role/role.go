package role

import (
	"errors"
	responseFactory "github.com/ilya-mezentsev/response-factory"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/domains/role/interfaces"
	"mini-roles-backend/source/domains/role/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"mini-roles-backend/source/domains/shared/services/validation"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

type Service struct {
	repository interfaces.RoleRepository
}

func New(repository interfaces.RoleRepository) Service {
	return Service{repository}
}

func (s Service) Create(request request.CreateRole) responseFactory.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	err := s.repository.Create(request.AccountId, request.Role)
	if err != nil {
		if errors.As(err, &sharedError.DuplicateUniqueKey{}) {
			return responseFactory.ClientError(sharedError.ServiceError{
				Code:        roleExistsCode,
				Description: roleExistsDescription,
			})
		}

		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to create role: %v", err)

		return response_factory.DefaultServerError()
	}

	return responseFactory.DefaultResponse()
}

func (s Service) RolesList(request request.RolesList) responseFactory.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	roles, err := s.repository.List(sharedSpec.AccountWithId{
		AccountId: request.AccountId,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to fetch roles from DB: %v", err)

		return response_factory.DefaultServerError()
	}

	return responseFactory.SuccessResponse(roles)
}

func (s Service) UpdateRole(request request.UpdateRole) responseFactory.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	err := s.repository.Update(request.AccountId, request.Role)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to update role: %v", err)

		return response_factory.DefaultServerError()
	}

	return responseFactory.DefaultResponse()
}

func (s Service) DeleteRole(request request.DeleteRole) responseFactory.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	err := s.repository.Delete(
		request.AccountId,
		request.RolesVersionId,
		request.RoleId,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to delete role: %v", err)

		return response_factory.DefaultServerError()
	}

	return responseFactory.DefaultResponse()
}
