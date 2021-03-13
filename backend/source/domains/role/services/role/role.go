package role

import (
	"mini-roles-backend/source/domains/role/interfaces"
	"mini-roles-backend/source/domains/role/request"
	shared "mini-roles-backend/source/domains/shared/interfaces"
	"mini-roles-backend/source/domains/shared/services/response_factory"
)

type Service struct {
	repository interfaces.RoleRepository
}

func New(repository interfaces.RoleRepository) Service {
	return Service{repository}
}

func (s Service) Create(request request.CreateRoleRequest) shared.HttpResponse {
	return response_factory.DefaultResponse()
}

func (s Service) RolesList(request request.RolesListRequest) shared.HttpResponse {
	return response_factory.DefaultResponse()
}

func (s Service) UpdateRole(request request.UpdateRoleRequest) shared.HttpResponse {
	return response_factory.DefaultResponse()
}

func (s Service) DeleteRole(request request.DeleteRoleRequest) shared.HttpResponse {
	return response_factory.DefaultResponse()
}
