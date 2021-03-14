package resource

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/domains/resource/interfaces"
	"mini-roles-backend/source/domains/resource/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	shared "mini-roles-backend/source/domains/shared/interfaces"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/hash"
	"mini-roles-backend/source/domains/shared/services/response_factory"
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

func (s Service) CreateResource(request request.CreateResourceRequest) shared.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	err = s.resourceRepository.Create(request.AccountId, request.Resource)
	if err != nil {
		if errors.As(err, &sharedError.DuplicateUniqueKey{}) {
			return response_factory.ClientError(sharedError.ServiceError{
				Code:        resourceExistsCode,
				Description: resourceExistsDescription,
			})
		}

		log.WithFields(log.Fields{
			"resource": request.Resource,
		}).Errorf("Unable to create resource in DB: %v", err)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	err = s.permissionRepository.AddResourcePermissions(
		request.AccountId,
		request.Resource.Id,
		s.generateResourcePermissions(request.Resource),
	)
	if err != nil {
		log.WithFields(log.Fields{
			"resource": request.Resource,
		}).Errorf("Unable to create permissions for new resource in DB: %v", err)

		// trying to delete resource (silently)
		_ = s.resourceRepository.Delete(request.AccountId, request.Resource.Id)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
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

func (s Service) ResourcesList(request request.ResourcesListRequest) shared.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	resources, err := s.resourceRepository.List(request.AccountId)
	if err != nil {
		log.Errorf("Unable to fetch resources from DB: %v", err)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	return response_factory.SuccessResponse(resources)
}

func (s Service) UpdateResource(request request.UpdateResourceRequest) shared.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	err = s.resourceRepository.Update(request.AccountId, request.Resource)
	if err != nil {
		log.WithFields(log.Fields{
			"resource": request.Resource,
		}).Errorf("Unable to fetch resources from DB: %v", err)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	return response_factory.DefaultResponse()
}

func (s Service) DeleteResource(request request.DeleteResourceRequest) shared.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	err = s.resourceRepository.Delete(request.AccountId, request.ResourceId)
	if err != nil {
		log.Errorf("Unable to update resource in DB: %v", err)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	return response_factory.DefaultResponse()
}
