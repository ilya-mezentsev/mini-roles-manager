package export

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"mini-roles-backend/source/domains/files/request"
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"mini-roles-backend/source/domains/shared/services/validation"
	sharedSpec "mini-roles-backend/source/domains/shared/spec"
	"os"
)

type Service struct {
	rolesFetcherRepository               sharedInterfaces.RolesFetcherRepository
	resourcesFetcherRepository           sharedInterfaces.ResourceFetcherRepository
	defaultRolesVersionFetcherRepository sharedInterfaces.DefaultRolesVersionFetcherRepository
}

func New(
	rolesFetcherRepository sharedInterfaces.RolesFetcherRepository,
	resourcesFetcherRepository sharedInterfaces.ResourceFetcherRepository,
	defaultRolesVersionFetcherRepository sharedInterfaces.DefaultRolesVersionFetcherRepository,
) Service {
	return Service{
		rolesFetcherRepository:               rolesFetcherRepository,
		resourcesFetcherRepository:           resourcesFetcherRepository,
		defaultRolesVersionFetcherRepository: defaultRolesVersionFetcherRepository,
	}
}

func (s Service) MakeExportFile(request request.ExportRequest) sharedInterfaces.Response {
	if validation.MakeErrorResponse(request) != nil {
		return response_factory.EmptyClientError()
	}

	spec := sharedSpec.AccountWithId{
		AccountId: request.AccountId,
	}
	resources, err := s.resourcesFetcherRepository.List(spec)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to fetch resources from DB: %v", err)

		return response_factory.EmptyServerError()
	}

	roles, err := s.rolesFetcherRepository.List(spec)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to fetch roles from DB: %v", err)

		return response_factory.EmptyServerError()
	}

	defaultRolesVersion, err := s.defaultRolesVersionFetcherRepository.Fetch(spec)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to fetch default roles version from db: %v", err)

		return response_factory.EmptyServerError()
	}

	jsonBytes, err := s.makeJSON(resources, roles, defaultRolesVersion)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to marshal data to json: %v", err)

		return response_factory.EmptyServerError()
	}

	tmpExportFilePath, err := s.createTmpFile(jsonBytes)
	if err != nil {
		log.WithFields(log.Fields{
			"request": request,
		}).Errorf("Unable to create tmp export file: %v", err)

		return response_factory.EmptyServerError()
	}

	return response_factory.SuccessResponse(tmpExportFilePath)
}

func (s Service) makeJSON(
	resources []sharedModels.Resource,
	roles []sharedModels.Role,
	defaultRolesVersion sharedModels.RolesVersion,
) ([]byte, error) {
	return json.Marshal(sharedModels.AppData{
		Resources:             resources,
		Roles:                 roles,
		DefaultRolesVersionId: defaultRolesVersion.Id,
	})
}

func (s Service) createTmpFile(content []byte) (string, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "export_json.*.json")
	if err != nil {
		return "", err
	}

	_, err = tmpFile.Write(content)
	if err != nil {
		return "", err
	}

	return tmpFile.Name(), tmpFile.Close()
}
