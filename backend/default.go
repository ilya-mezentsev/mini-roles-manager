package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	"mini-roles-backend/source/db/schema"
	accountInfoRepositoryConstructor "mini-roles-backend/source/domains/account/repositories/info"
	registrationRepositoryConstructor "mini-roles-backend/source/domains/account/repositories/registration"
	sessionRepositoryConstructor "mini-roles-backend/source/domains/account/repositories/session"
	"mini-roles-backend/source/domains/account/services/info"
	"mini-roles-backend/source/domains/account/services/registration"
	"mini-roles-backend/source/domains/account/services/session"
	"mini-roles-backend/source/domains/account/services/session_check"
	"mini-roles-backend/source/domains/files/services/export"
	permissionListRepositoryConstructor "mini-roles-backend/source/domains/permission/repositories/permission_db"
	"mini-roles-backend/source/domains/permission/services/permission"
	permissionCreatorRepositoryConstructor "mini-roles-backend/source/domains/resource/repositories/permission"
	resourceRepositoryConstructor "mini-roles-backend/source/domains/resource/repositories/resource"
	"mini-roles-backend/source/domains/resource/services/resource"
	roleRepositoryConstructor "mini-roles-backend/source/domains/role/repositories/role"
	"mini-roles-backend/source/domains/role/services/role"
	"mini-roles-backend/source/entrypoints/web"
)

var configFilePath = flag.String("config", "/dev/null", "Set path to configs file")

func fullInit(r *gin.Engine) int {
	configsRepository := config.MustNew(*configFilePath)

	db := connection.MustGetConnection(configsRepository)
	db.MustExec(schema.Schema)

	registrationRepository := registrationRepositoryConstructor.New(db)
	sessionRepository := sessionRepositoryConstructor.New(db)
	accountInfoRepository := accountInfoRepositoryConstructor.New(db)
	permissionListRepository := permissionListRepositoryConstructor.New(db)
	permissionCreatorRepository := permissionCreatorRepositoryConstructor.New(db)
	resourceRepository := resourceRepositoryConstructor.New(db)
	roleRepository := roleRepositoryConstructor.New(db)

	registrationService := registration.New(registrationRepository)
	sessionService := session.New(sessionRepository, configsRepository)
	accountInfoService := info.New(accountInfoRepository)
	sessionCheckService := session_check.New(sessionRepository)
	permissionService := permission.New(permissionListRepository)
	resourceService := resource.New(resourceRepository, permissionCreatorRepository)
	roleService := role.New(roleRepository)
	exportService := export.New(roleRepository, resourceRepository)

	web.FullInit(
		r,

		registrationService,
		sessionService,
		sessionCheckService,
		accountInfoService,

		permissionService,

		resourceService,

		roleService,

		exportService,
	)

	return configsRepository.ServerPort()
}
