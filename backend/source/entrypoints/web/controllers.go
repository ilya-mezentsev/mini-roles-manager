package web

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/account/services/info"
	"mini-roles-backend/source/domains/account/services/registration"
	"mini-roles-backend/source/domains/account/services/session"
	"mini-roles-backend/source/domains/account/services/session_check"
	"mini-roles-backend/source/domains/files/services/export"
	"mini-roles-backend/source/domains/permission/services/permission"
	"mini-roles-backend/source/domains/resource/services/resource"
	"mini-roles-backend/source/domains/role/services/role"
	"mini-roles-backend/source/domains/role/services/roles_version"
	"mini-roles-backend/source/entrypoints/web/controllers/account"
	filesControllerConstructor "mini-roles-backend/source/entrypoints/web/controllers/files"
	permissionControllerConstructor "mini-roles-backend/source/entrypoints/web/controllers/permission"
	resourceControllerConstructor "mini-roles-backend/source/entrypoints/web/controllers/resource"
	roleControllerConstructor "mini-roles-backend/source/entrypoints/web/controllers/role"
	rolesVersionControllerConstructor "mini-roles-backend/source/entrypoints/web/controllers/roles_version"
	"mini-roles-backend/source/entrypoints/web/middleware/cookie"
	"mini-roles-backend/source/entrypoints/web/middleware/header"
)

func FullInit(
	r *gin.Engine,

	registrationService registration.Service,
	sessionService session.Service,
	sessionCheckService session_check.Service,
	accountInfoService info.Service,

	permissionService permission.Service,

	resourceService resource.Service,

	rolesVersionService roles_version.Service,

	roleService role.Service,

	exportService export.Service,

	sharedMiddlewares ...gin.HandlerFunc,
) {
	checkCookieMiddleware := cookie.New(sessionCheckService)
	checkHeaderMiddleware := header.New(sessionCheckService)

	accountController := account.New(registrationService, sessionService, accountInfoService)
	permissionController := permissionControllerConstructor.New(permissionService)
	resourceController := resourceControllerConstructor.New(resourceService)
	rolesVersionController := rolesVersionControllerConstructor.New(rolesVersionService)
	roleController := roleControllerConstructor.New(roleService)
	filesController := filesControllerConstructor.New(exportService)

	webAppGroup := r.Group("/web-app")
	webAppGroup.Use(sharedMiddlewares...)

	webAppGroup.POST("/registration/user", accountController.Register)

	webAppGroup.GET("/session", accountController.Login)
	webAppGroup.POST("/session", accountController.SignIn)
	webAppGroup.DELETE("/session", accountController.SignOut)

	cookieTokenAuthorizedGroup := webAppGroup.Group("/")
	cookieTokenAuthorizedGroup.Use(checkCookieMiddleware.HasSessionInCookie())
	{
		cookieTokenAuthorizedGroup.GET("/account/info", accountController.GetAccountInfo)
		cookieTokenAuthorizedGroup.PATCH("/account/credentials", accountController.UpdateCredentials)

		cookieTokenAuthorizedGroup.GET("/permissions", permissionController.ResolveResourceAccessEffect)

		cookieTokenAuthorizedGroup.GET("/resources", resourceController.ResourcesList)
		cookieTokenAuthorizedGroup.POST("/resource", resourceController.CreateResource)
		cookieTokenAuthorizedGroup.PATCH("/resource", resourceController.UpdateResource)
		cookieTokenAuthorizedGroup.DELETE("/resource/:resource_id", resourceController.DeleteResource)

		cookieTokenAuthorizedGroup.GET("/roles-versions", rolesVersionController.RolesVersionsList)
		cookieTokenAuthorizedGroup.POST("/roles-version", rolesVersionController.CreateRolesVersion)
		cookieTokenAuthorizedGroup.PATCH("/roles-version", rolesVersionController.UpdateRolesVersion)
		cookieTokenAuthorizedGroup.DELETE("/roles-version/:roles_version_id", rolesVersionController.DeleteRolesVersion)

		cookieTokenAuthorizedGroup.GET("/roles", roleController.RolesList)
		cookieTokenAuthorizedGroup.POST("/role", roleController.CreateRole)
		cookieTokenAuthorizedGroup.PATCH("/role", roleController.UpdateRole)
		cookieTokenAuthorizedGroup.DELETE("/role/:roles_version_id/:role_id", roleController.DeleteRole)

		cookieTokenAuthorizedGroup.GET("/app-data/export", filesController.Export)
	}

	headerTokenAuthorized := r.Group("/public")
	headerTokenAuthorized.Use(
		append(
			sharedMiddlewares,
			checkHeaderMiddleware.HasSessionInHeader(),
		)...,
	)
	{
		headerTokenAuthorized.GET("/permissions", permissionController.ResolveResourceAccessEffect)
	}
}

func MinimalInit(
	r *gin.Engine,

	permissionService permission.Service,

	sharedMiddlewares ...gin.HandlerFunc,
) {
	permissionController := permissionControllerConstructor.New(permissionService)

	r.Use(sharedMiddlewares...)

	r.GET("/permissions", permissionController.ResolveResourceAccessEffect)
}
