package roles_version

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
	repository interfaces.RolesVersionRepository
}

func New(repository interfaces.RolesVersionRepository) Service {
	return Service{repository}
}

func (s Service) CreateRolesVersion(request request.CreateRolesVersion) responseFactory.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	err := s.repository.Create(request.AccountId, request.RolesVersion)
	if err != nil {
		if errors.As(err, &sharedError.DuplicateUniqueKey{}) {
			return responseFactory.ClientError(sharedError.ServiceError{
				Code:        rolesVersionExistsCode,
				Description: rolesVersionExistsDescription,
			})
		}

		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to create role: %v", err)

		return response_factory.DefaultServerError()
	}

	return responseFactory.DefaultResponse()
}

func (s Service) RolesVersionList(request request.RolesVersionList) responseFactory.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	rolesVersionList, err := s.repository.List(sharedSpec.AccountWithId{
		AccountId: request.AccountId,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to fetch roles versions from DB: %v", err)

		return response_factory.DefaultServerError()
	}

	return responseFactory.SuccessResponse(rolesVersionList)
}

func (s Service) UpdateRolesVersion(request request.UpdateRolesVersion) responseFactory.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	err := s.repository.Update(request.AccountId, request.RolesVersion)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to update roles version: %v", err)

		return response_factory.DefaultServerError()
	}

	return responseFactory.DefaultResponse()
}

func (s Service) DeleteRolesVersion(request request.DeleteRolesVersion) responseFactory.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	rolesVersionList, err := s.repository.List(sharedSpec.AccountWithId{
		AccountId: request.AccountId,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to fetch roles versions from DB: %v", err)

		return response_factory.DefaultServerError()
	}

	if len(rolesVersionList) < 2 {
		return responseFactory.ClientError(sharedError.ServiceError{
			Code:        cannotDeleteLastRolesVersionCode,
			Description: cannotDeleteLastRolesVersionDescription,
		})
	}

	err = s.repository.Delete(request.AccountId, request.RolesVersionId)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to delete roles version: %v", err)

		return response_factory.DefaultServerError()
	}

	return responseFactory.DefaultResponse()
}
