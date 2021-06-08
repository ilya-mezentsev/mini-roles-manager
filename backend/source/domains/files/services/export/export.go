package export

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"mini-roles-backend/source/domains/files/models"
	"mini-roles-backend/source/domains/files/request"
	sharedInterfaces "mini-roles-backend/source/domains/shared/interfaces"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	"mini-roles-backend/source/domains/shared/services/response_factory"
	"mini-roles-backend/source/domains/shared/services/validation"
	"os"
)

type Service struct {
	rolesFetcherRepository     sharedInterfaces.RolesFetcherRepository
	resourcesFetcherRepository sharedInterfaces.ResourceFetcherRepository
}

func New(
	rolesFetcherRepository sharedInterfaces.RolesFetcherRepository,
	resourcesFetcherRepository sharedInterfaces.ResourceFetcherRepository,
) Service {
	return Service{
		rolesFetcherRepository:     rolesFetcherRepository,
		resourcesFetcherRepository: resourcesFetcherRepository,
	}
}

func (s Service) MakeExportFile(request request.ExportRequest) sharedInterfaces.Response {
	if validation.MakeErrorResponse(request) != nil {
		return response_factory.EmptyClientError()
	}

	resources, err := s.resourcesFetcherRepository.List(request.AccountId)
	if err != nil {
		log.Errorf("Unable to fetch resources from DB: %v", err)

		return response_factory.EmptyServerError()
	}

	roles, err := s.rolesFetcherRepository.List(request.AccountId)
	if err != nil {
		log.Errorf("Unable to fetch roles from DB: %v", err)

		return response_factory.EmptyServerError()
	}

	jsonBytes, err := s.makeJSON(resources, roles)
	if err != nil {
		log.Errorf("Unable to marshal data to json: %v", err)

		return response_factory.EmptyServerError()
	}

	tmpExportFilePath, err := s.createTmpFile(jsonBytes)
	if err != nil {
		log.Errorf("Unable to create tmp export file: %v", err)

		return response_factory.EmptyServerError()
	}

	return response_factory.SuccessResponse(tmpExportFilePath)
}

func (s Service) makeJSON(
	resources []sharedModels.Resource,
	roles []sharedModels.Role,
) ([]byte, error) {
	return json.Marshal(models.JSONRepresentation{
		Resources: resources,
		Roles:     roles,
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
