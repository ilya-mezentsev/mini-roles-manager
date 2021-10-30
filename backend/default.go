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
	resetResourcesRepositoryConstructor "mini-roles-backend/source/domains/files/repositories/reset_resources"
	resetRolesRepositoryConstructor "mini-roles-backend/source/domains/files/repositories/reset_roles"
	resetRolesVersionRepositoryConstructor "mini-roles-backend/source/domains/files/repositories/reset_roles_version"
	"mini-roles-backend/source/domains/files/services/export"
	"mini-roles-backend/source/domains/files/services/import_file"
	permissionCacheConstructor "mini-roles-backend/source/domains/permission/cache/permission"
	defaultRolesVersionRepositoryConstructor "mini-roles-backend/source/domains/permission/repositories/db/default_roles_version"
	permissionRepositoryConstructor "mini-roles-backend/source/domains/permission/repositories/db/permission"
	"mini-roles-backend/source/domains/permission/services/permission"
	permissionCreatorRepositoryConstructor "mini-roles-backend/source/domains/resource/repositories/permission"
	resourceRepositoryConstructor "mini-roles-backend/source/domains/resource/repositories/resource"
	"mini-roles-backend/source/domains/resource/services/resource"
	roleRepositoryConstructor "mini-roles-backend/source/domains/role/repositories/role"
	rolesVersionRepositoryConstructor "mini-roles-backend/source/domains/role/repositories/roles_version"
	"mini-roles-backend/source/domains/role/services/role"
	"mini-roles-backend/source/domains/role/services/roles_version"
	"mini-roles-backend/source/entrypoints/web"
)

var configFilePath = flag.String("config", "/dev/null", "Set path to configs file")

func fullInit(
	r *gin.Engine,
	sharedMiddlewares ...gin.HandlerFunc,
) int {
	configsRepository := config.MustNew(*configFilePath)

	db := connection.MustGetConnection(configsRepository)
	db.MustExec(schema.Schema)

	registrationRepository := registrationRepositoryConstructor.New(db)
	sessionRepository := sessionRepositoryConstructor.New(db)
	accountInfoRepository := accountInfoRepositoryConstructor.New(db)
	permissionListRepository := permissionRepositoryConstructor.New(db)
	permissionCreatorRepository := permissionCreatorRepositoryConstructor.New(db)
	resourceRepository := resourceRepositoryConstructor.New(db)
	rolesVersionRepository := rolesVersionRepositoryConstructor.New(db)
	defaultRolesVersionRepository := defaultRolesVersionRepositoryConstructor.New(db)
	roleRepository := roleRepositoryConstructor.New(db)
	resetRolesVersionRepository := resetRolesVersionRepositoryConstructor.New(db)
	resetResourcesRepository := resetResourcesRepositoryConstructor.New(db)
	resetRolesRepository := resetRolesRepositoryConstructor.New(db)

	permissionCache := permissionCacheConstructor.New(
		configsRepository.CachePermissionLifetime(),
		permissionListRepository,
	)

	registrationService := registration.New(registrationRepository, rolesVersionRepository)
	sessionService := session.New(sessionRepository, configsRepository)
	accountInfoService := info.New(accountInfoRepository)
	sessionCheckService := session_check.New(sessionRepository)
	permissionService := permission.New(permissionCache, defaultRolesVersionRepository)
	resourceService := resource.New(
		resourceRepository,
		permissionCreatorRepository,
		permissionCache,
	)
	rolesVersionService := roles_version.New(rolesVersionRepository, permissionCache)
	roleService := role.New(roleRepository, permissionCache)
	exportService := export.New(
		roleRepository,
		resourceRepository,
		defaultRolesVersionRepository,
	)
	importService := import_file.New(
		resetRolesVersionRepository,
		resetResourcesRepository,
		resetRolesRepository,
	)

	web.FullInit(
		r,

		registrationService,
		sessionService,
		sessionCheckService,
		accountInfoService,

		permissionService,

		resourceService,

		rolesVersionService,

		roleService,

		exportService,
		importService,

		sharedMiddlewares...,
	)

	return configsRepository.ServerPort()
}
