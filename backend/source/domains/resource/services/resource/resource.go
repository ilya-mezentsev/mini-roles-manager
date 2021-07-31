package resource

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/domains/resource/interfaces"
	"mini-roles-backend/source/domains/resource/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/hash"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"mini-roles-backend/source/domains/shared/services/validation"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

type Service struct {
	resourceRepository   interfaces.ResourceRepository
	permissionRepository interfaces.PermissionRepository
}

func New(
	resourceRepository interfaces.ResourceRepository,
	permissionRepository interfaces.PermissionRepository,
) Service {
	return Service{
		resourceRepository:   resourceRepository,
		permissionRepository: permissionRepository,
	}
}

func (s Service) CreateResource(request request.CreateResource) sharedInterfaces.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	err := s.resourceRepository.Create(request.AccountId, request.Resource)
	if err != nil {
		if errors.As(err, &sharedError.DuplicateUniqueKey{}) {
			return response_factory.ClientError(sharedError.ServiceError{
				Code:        resourceExistsCode,
				Description: resourceExistsDescription,
			})
		}

		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to create resource in DB: %v", err)

		return response_factory.DefaultServerError()
	}

	err = s.permissionRepository.AddResourcePermissions(
		request.AccountId,
		request.Resource.Id,
		s.generateResourcePermissions(request.Resource),
	)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to create permissions for new resource in DB: %v", err)

		// if resource deleting failed - this log will help to find it and delete manually
		log.Warningf(
			"Got error while deleting resource (id = %s): %v",
			s.resourceRepository.Delete(request.AccountId, request.Resource.Id),
			request.Resource.Id,
		)

		return response_factory.DefaultServerError()
	}

	return response_factory.DefaultResponse()
}

func (s Service) generateResourcePermissions(resource sharedModels.Resource) []sharedModels.Permission {
	var permissions []sharedModels.Permission
	for _, operation := range resourcesOperations {
		for _, operationEffect := range resourcesOperationsEffects {
			permissionId := hash.Md5WithTimeAsKey(fmt.Sprintf(
				"%s:%s:%s",
				resource.Id,
				operation,
				operationEffect,
			))

			permissions = append(permissions, sharedModels.Permission{
				Id:        sharedModels.PermissionId(permissionId),
				Resource:  resource,
				Operation: operation,
				Effect:    operationEffect,
			})
		}
	}

	return permissions
}

func (s Service) ResourcesList(request request.ResourcesList) sharedInterfaces.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	resources, err := s.resourceRepository.List(sharedSpec.AccountWithId{
		AccountId: request.AccountId,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to fetch resources from DB: %v", err)

		return response_factory.DefaultServerError()
	}

	return response_factory.SuccessResponse(resources)
}

func (s Service) UpdateResource(request request.UpdateResource) sharedInterfaces.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	err := s.resourceRepository.Update(request.AccountId, request.Resource)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to fetch resources from DB: %v", err)

		return response_factory.DefaultServerError()
	}

	return response_factory.DefaultResponse()
}

func (s Service) DeleteResource(request request.DeleteResource) sharedInterfaces.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	err := s.resourceRepository.Delete(request.AccountId, request.ResourceId)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to update resource in DB: %v", err)

		return response_factory.DefaultServerError()
	}

	return response_factory.DefaultResponse()
}
