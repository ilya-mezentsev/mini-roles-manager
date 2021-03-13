package permission

import (
	"mini-roles-backend/source/domains/permission/interfaces"
)

type Service struct {
	repository interfaces.PermissionRepository
}

func New(repository interfaces.PermissionRepository) Service {
	return Service{repository}
}
