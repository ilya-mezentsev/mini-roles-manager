package import_file

import (
	responseFactory "github.com/ilya-mezentsev/response-factory"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"mini-roles-backend/source/domains/files/interfaces"
	"mini-roles-backend/source/domains/files/request"
	"mini-roles-backend/source/domains/files/services/validation"
	sharedError "mini-roles-backend/source/domains/shared/error"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"strings"
)

type Service struct {
	resetRolesVersionRepository interfaces.ResetRolesVersionRepository
	resetResourcesRepository    interfaces.ResetResourcesRepository
	resetRolesRepository        interfaces.ResetRolesRepository
}

func New(
	resetRolesVersionRepository interfaces.ResetRolesVersionRepository,
	resetResourcesRepository interfaces.ResetResourcesRepository,
	resetRolesRepository interfaces.ResetRolesRepository,
) Service {
	return Service{
		resetRolesVersionRepository: resetRolesVersionRepository,
		resetResourcesRepository:    resetResourcesRepository,
		resetRolesRepository:        resetRolesRepository,
	}
}

func (s Service) ImportFromFile(request request.ImportRequest) responseFactory.Response {
	appDataBytes, err := ioutil.ReadAll(request.File)
	if err != nil {
		log.Errorf("Unable to read file: %v", err)

		return response_factory.DefaultServerError()
	}

	appData, errors := validation.Validate(appDataBytes)
	if len(errors) > 0 {
		return responseFactory.ClientError(sharedError.ServiceError{
			Code:        invalidImportFileCode,
			Description: strings.Join(errors, " | "),
		})
	}

	err = s.resetRolesVersionRepository.Reset(
		request.AccountId,
		sharedModels.RolesVersion{
			Id: appData.DefaultRolesVersionId,
		},
		s.makeRolesVersions(appData),
	)
	if err != nil {
		log.WithFields(log.Fields{
			"account_id": request.AccountId,
			"appData":    appData,
		}).Errorf("Unable to reset roles versions: %v", err)

		return response_factory.DefaultServerError()
	}

	err = s.resetResourcesRepository.Reset(
		request.AccountId,
		appData.Resources,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"account_id": request.AccountId,
			"appData":    appData,
		}).Errorf("Unable to reset resources: %v", err)

		return response_factory.DefaultServerError()
	}

	err = s.resetRolesRepository.Reset(
		request.AccountId,
		appData.Roles,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"account_id": request.AccountId,
			"appData":    appData,
		}).Errorf("Unable to reset roles: %v", err)

		return response_factory.DefaultServerError()
	}

	return responseFactory.DefaultResponse()
}

func (s Service) makeRolesVersions(appData sharedModels.AppData) []sharedModels.RolesVersion {
	var rolesVersions []sharedModels.RolesVersion
	uniqueRolesVersion := map[sharedModels.RolesVersionId]int{}
	for _, role := range appData.Roles {
		_, versionFound := uniqueRolesVersion[role.VersionId]
		if versionFound || role.VersionId == appData.DefaultRolesVersionId {
			continue
		}

		uniqueRolesVersion[role.VersionId] = 0
		rolesVersions = append(rolesVersions, sharedModels.RolesVersion{
			Id: role.VersionId,
		})
	}

	return rolesVersions
}
