package web

import (
	"github.com/gin-gonic/gin"
	"mini-roles-backend/source/domains/account/services/info"
	"mini-roles-backend/source/domains/account/services/registration"
	"mini-roles-backend/source/domains/account/services/session"
	"mini-roles-backend/source/domains/account/services/session_check"
	"mini-roles-backend/source/domains/permission/services/permission"
	"mini-roles-backend/source/domains/resource/services/resource"
	"mini-roles-backend/source/domains/role/services/role"
	"mini-roles-backend/source/entrypoints/web/controllers/account"
	permissionControllerConstructor "mini-roles-backend/source/entrypoints/web/controllers/permission"
	resourceControllerConstructor "mini-roles-backend/source/entrypoints/web/controllers/resource"
	roleControllerConstructor "mini-roles-backend/source/entrypoints/web/controllers/role"
	"mini-roles-backend/source/entrypoints/web/middleware/cookie"
	"mini-roles-backend/source/entrypoints/web/middleware/header"
)

func Init(
	r *gin.Engine,

	registrationService registration.Service,
	sessionService session.Service,
	sessionCheckService session_check.Service,
	accountInfoService info.Service,

	permissionService permission.Service,

	resourceService resource.Service,

	roleService role.Service,
) {
	checkCookieMiddleware := cookie.New(sessionCheckService)
	checkHeaderMiddleware := header.New(sessionCheckService)

	accountController := account.New(registrationService, sessionService, accountInfoService)
	permissionController := permissionControllerConstructor.New(permissionService)
	resourceController := resourceControllerConstructor.New(resourceService)
	roleController := roleControllerConstructor.New(roleService)

	r.POST("/registration/user", accountController.Register)

	r.GET("/session", accountController.Login)
	r.POST("/session", accountController.SignIn)
	r.DELETE("/session", accountController.SignOut)

	cookieTokenAuthorized := r.Group("/")
	cookieTokenAuthorized.Use(checkCookieMiddleware.HasSessionInCookie())
	{
		cookieTokenAuthorized.GET("/account/info", accountController.GetAccountInfo)
		cookieTokenAuthorized.PATCH("/account/credentials", accountController.UpdateCredentials)

		cookieTokenAuthorized.POST("/check-permissions", permissionController.ResolveResourceAccessEffect)

		cookieTokenAuthorized.GET("/resources", resourceController.ResourcesList)
		cookieTokenAuthorized.POST("/resource", resourceController.CreateResource)
		cookieTokenAuthorized.PATCH("/resource", resourceController.UpdateResource)
		cookieTokenAuthorized.DELETE("/resource/:resource_id", resourceController.DeleteResource)

		cookieTokenAuthorized.GET("/roles", roleController.RolesList)
		cookieTokenAuthorized.POST("/role", roleController.CreateRole)
		cookieTokenAuthorized.PATCH("/role", roleController.UpdateRole)
		cookieTokenAuthorized.DELETE("/role/:role_id", roleController.DeleteRole)
	}

	headerTokenAuthorized := r.Group("/")
	headerTokenAuthorized.Use(checkHeaderMiddleware.HasSessionInHeader())
	{
		headerTokenAuthorized.POST("/permissions", permissionController.ResolveResourceAccessEffect)
	}
}
