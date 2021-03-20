package permission

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/domains/permission/interfaces"
	"mini-roles-backend/source/domains/permission/models"
	"mini-roles-backend/source/domains/permission/request"
	sharedError "mini-roles-backend/source/domains/shared/error"
	shared "mini-roles-backend/source/domains/shared/interfaces"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/response_factory"
)

type Service struct {
	repository interfaces.PermissionRepository
}

func New(repository interfaces.PermissionRepository) Service {
	return Service{repository}
}

func (s Service) HasPermission(request request.PermissionAccess) shared.Response {
	err := validator.New().Struct(request)
	if err != nil {
		return response_factory.ClientError(sharedError.ServiceError{
			Code:        sharedError.ValidationErrorCode,
			Description: err.Error(),
		})
	}

	permissions, err := s.repository.List(request.AccountId, request.RoleId)
	if err != nil {
		log.Errorf("Unable to fetch permissions from DB: %v", err)

		return response_factory.ServerError(sharedError.ServiceError{
			Code:        sharedError.ServerErrorCode,
			Description: sharedError.ServerErrorDescription,
		})
	}

	return s.makePermissionsResponse(request, permissions)
}

func (s Service) makePermissionsResponse(
	request request.PermissionAccess,
	rolePermissions []sharedModels.Permission,
) shared.Response {
	effect := s.effectForRequestedRole(request, rolePermissions)
	if effect.IsEmpty() {
		return s.checkPermissionsForLinkingResources(request, rolePermissions)
	} else {
		return s.makeResponseFromEffect(effect)
	}
}

func (s Service) effectForRequestedRole(
	request request.PermissionAccess,
	rolePermissions []sharedModels.Permission,
) interfaces.Effect {
	for _, permission := range rolePermissions {
		if permission.Resource.Id == request.ResourceId &&
			permission.Operation == request.Operation {
			if permission.Effect == permitEffectCode {
				return models.PermitEffect{}
			} else {
				return models.DenyEffect{}
			}
		}
	}

	return models.MissedEffect{}
}

func (s Service) makeResponseFromEffect(effect interfaces.Effect) shared.Response {
	if effect.IsPermit() {
		return response_factory.SuccessResponse(models.EffectResponse{
			Effect: permitEffectCode,
		})
	} else {
		return response_factory.SuccessResponse(models.EffectResponse{
			Effect: denyEffectCode,
		})
	}
}

func (s Service) checkPermissionsForLinkingResources(
	request request.PermissionAccess,
	rolePermissions []sharedModels.Permission,
) shared.Response {
	linkingResource, resourceFound := s.findLinkingResources(request.ResourceId, rolePermissions)
	if resourceFound {
		return s.makePermissionsResponse(
			request.WithResourceId(linkingResource.Id),
			rolePermissions,
		)
	} else {
		return s.makeResponseFromEffect(models.DenyEffect{})
	}
}

func (s Service) findLinkingResources(
	requestedResourceId sharedModels.ResourceId,
	rolePermissions []sharedModels.Permission,
) (foundResource sharedModels.Resource, resourceFound bool) {
	for _, permission := range rolePermissions {
		for _, resourceId := range permission.Resource.LinksTo {
			if resourceId == requestedResourceId {
				resourceFound = true
				foundResource = permission.Resource
			}
		}
	}

	return
}
