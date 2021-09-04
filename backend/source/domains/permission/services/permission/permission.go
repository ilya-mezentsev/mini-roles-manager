package permission

import (
	responseFactory "github.com/ilya-mezentsev/response-factory"
	log "github.com/sirupsen/logrus"
	"mini-roles-backend/source/domains/permission/interfaces"
	"mini-roles-backend/source/domains/permission/models"
	"mini-roles-backend/source/domains/permission/request"
	"mini-roles-backend/source/domains/permission/spec"
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"mini-roles-backend/source/domains/shared/services/validation"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
)

type Service struct {
	permissionRepository          interfaces.PermissionRepository
	defaultRolesVersionRepository sharedInterfaces.DefaultRolesVersionFetcherRepository
}

func New(
	permissionRepository interfaces.PermissionRepository,
	defaultRolesVersionRepository sharedInterfaces.DefaultRolesVersionFetcherRepository,
) Service {
	return Service{
		permissionRepository:          permissionRepository,
		defaultRolesVersionRepository: defaultRolesVersionRepository,
	}
}

func (s Service) HasPermission(
	accountId sharedModels.AccountId,
	request request.PermissionAccess,
) responseFactory.Response {
	invalidRequestResponse := validation.MakeErrorResponse(request)
	if invalidRequestResponse != nil {
		return invalidRequestResponse
	}

	rolesVersionId, err := s.rolesVersionId(
		accountId,
		request,
	)
	if err != nil {
		return response_factory.DefaultServerError()
	}

	permissions, err := s.permissionRepository.List(spec.PermissionWithAccountIdAndRoleId{
		AccountId:      accountId,
		RoleId:         request.RoleId,
		RolesVersionId: rolesVersionId,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"account_id": accountId,
			"request":    request,
		}).Errorf("Unable to fetch permissions from DB: %v", err)

		return response_factory.DefaultServerError()
	}

	return s.makePermissionsResponse(request, permissions)
}

func (s Service) rolesVersionId(
	accountId sharedModels.AccountId,
	request request.PermissionAccess,
) (sharedModels.RolesVersionId, error) {
	if request.RolesVersionId != "" {
		return request.RolesVersionId, nil
	}

	defaultRolesVersion, err := s.defaultRolesVersionRepository.Fetch(sharedSpec.AccountWithId{
		AccountId: accountId,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"account_id": accountId,
			"request":    request,
		}).Errorf("Unable to fetch default roles version from DB: %v", err)
	}

	return defaultRolesVersion.Id, err
}

func (s Service) makePermissionsResponse(
	request request.PermissionAccess,
	rolePermissions []sharedModels.Permission,
) responseFactory.Response {
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

func (s Service) makeResponseFromEffect(effect interfaces.Effect) responseFactory.Response {
	if effect.IsPermit() {
		return responseFactory.SuccessResponse(models.EffectResponse{
			Effect: permitEffectCode,
		})
	} else {
		return responseFactory.SuccessResponse(models.EffectResponse{
			Effect: denyEffectCode,
		})
	}
}

func (s Service) checkPermissionsForLinkingResources(
	request request.PermissionAccess,
	rolePermissions []sharedModels.Permission,
) responseFactory.Response {
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
