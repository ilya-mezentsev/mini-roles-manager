package web

import (
	"github.com/gin-gonic/gin"
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

	permissionService permission.Service,

	resourceService resource.Service,

	roleService role.Service,
) {
	checkCookieMiddleware := cookie.New(sessionCheckService)
	checkHeaderMiddleware := header.New(sessionCheckService)

	accountController := account.New(registrationService, sessionService)
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
		r.GET("/resources", resourceController.ResourcesList)
		r.POST("/resource", resourceController.CreateResource)
		r.PATCH("/resource", resourceController.UpdateResource)
		r.DELETE("/resource/:resource_id", resourceController.DeleteResource)

		r.GET("/roles", roleController.RolesList)
		r.POST("/role", roleController.CreateRole)
		r.PATCH("/role", roleController.UpdateRole)
		r.DELETE("/role/:role_id", roleController.DeleteRole)
	}

	headerTokenAuthorized := r.Group("/")
	headerTokenAuthorized.Use(checkHeaderMiddleware.HasSessionInHeader())
	{
		r.POST("/permissions", permissionController.ResolveResourceAccessEffect)
	}
}
