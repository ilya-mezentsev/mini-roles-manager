package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mini-roles-backend/source/config"
	"mini-roles-backend/source/db/connection"
	"mini-roles-backend/source/db/schema"
	accountInterfaces "mini-roles-backend/source/domains/account/interfaces"
	accountInfoRepositoryConstructor "mini-roles-backend/source/domains/account/repositories/info"
	registrationRepositoryConstructor "mini-roles-backend/source/domains/account/repositories/registration"
	sessionRepositoryConstructor "mini-roles-backend/source/domains/account/repositories/session"
	"mini-roles-backend/source/domains/account/services/info"
	"mini-roles-backend/source/domains/account/services/registration"
	"mini-roles-backend/source/domains/account/services/session"
	"mini-roles-backend/source/domains/account/services/session_check"
	permissionInterfaces "mini-roles-backend/source/domains/permission/interfaces"
	permissionListRepositoryConstructor "mini-roles-backend/source/domains/permission/repositories/permission"
	"mini-roles-backend/source/domains/permission/services/permission"
	resourceInterfaces "mini-roles-backend/source/domains/resource/interfaces"
	permissionCreatorRepositoryConstructor "mini-roles-backend/source/domains/resource/repositories/permission"
	resourceRepositoryConstructor "mini-roles-backend/source/domains/resource/repositories/resource"
	"mini-roles-backend/source/domains/resource/services/resource"
	roleInterfaces "mini-roles-backend/source/domains/role/interfaces"
	roleRepositoryConstructor "mini-roles-backend/source/domains/role/repositories/role"
	"mini-roles-backend/source/domains/role/services/role"
	"mini-roles-backend/source/entrypoints/web"
)

var (
	configFilePath    = flag.String("config", "/dev/null", "Set path to configs file")
	configsRepository config.Repository

	registrationRepository      accountInterfaces.RegistrationRepository
	sessionRepository           accountInterfaces.SessionRepository
	accountInfoRepository       accountInfoRepositoryConstructor.Repository
	permissionListRepository    permissionInterfaces.PermissionRepository
	permissionCreatorRepository resourceInterfaces.PermissionRepository
	resourceRepository          resourceInterfaces.ResourceRepository
	roleRepository              roleInterfaces.RoleRepository

	registrationService registration.Service
	sessionService      session.Service
	accountInfoService  info.Service
	sessionCheckService session_check.Service
	permissionService   permission.Service
	resourceService     resource.Service
	roleService         role.Service
)

func init() {
	flag.Parse()

	configsRepository = config.MustNew(*configFilePath)

	db := connection.MustGetConnection(configsRepository)
	db.MustExec(schema.Schema)

	registrationRepository = registrationRepositoryConstructor.New(db)
	sessionRepository = sessionRepositoryConstructor.New(db)
	accountInfoRepository = accountInfoRepositoryConstructor.New(db)
	permissionListRepository = permissionListRepositoryConstructor.New(db)
	permissionCreatorRepository = permissionCreatorRepositoryConstructor.New(db)
	resourceRepository = resourceRepositoryConstructor.New(db)
	roleRepository = roleRepositoryConstructor.New(db)

	registrationService = registration.New(registrationRepository)
	sessionService = session.New(sessionRepository, configsRepository)
	accountInfoService = info.New(accountInfoRepository)
	sessionCheckService = session_check.New(sessionRepository)
	permissionService = permission.New(permissionListRepository)
	resourceService = resource.New(resourceRepository, permissionCreatorRepository)
	roleService = role.New(roleRepository)
}

func main() {
	r := gin.Default()

	web.Init(
		r,

		registrationService,
		sessionService,
		sessionCheckService,
		accountInfoService,

		permissionService,

		resourceService,

		roleService,
	)

	err := r.Run(fmt.Sprintf(":%d", configsRepository.ServerPort()))
	if err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}
}
