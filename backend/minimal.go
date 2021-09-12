package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"mini-roles-backend/source/domains/files/services/validation"
	defaultRolesVersionRepositoryConstructor "mini-roles-backend/source/domains/permission/repositories/memory/default_roles_version"
	permissionRepositoryConstructor "mini-roles-backend/source/domains/permission/repositories/memory/permission"
	"mini-roles-backend/source/domains/permission/services/permission"
	"mini-roles-backend/source/entrypoints/web"
	"strings"
)

var (
	appDataPath = flag.String(
		"app-data",
		"/dev/null",
		"Set path to file with application data (resources, roles, etc)",
	)
	serverPort = flag.Int(
		"port",
		8080,
		"Server port",
	)
)

func minimalInit(
	r *gin.Engine,
	sharedMiddlewares ...gin.HandlerFunc,
) int {
	appDataBytes, err := ioutil.ReadFile(*appDataPath)
	if err != nil {
		log.Fatalf("Unable to read app data file by path %s: %v", *appDataPath, err)
	}

	appData, validationErrors := validation.Validate(appDataBytes)
	if len(validationErrors) > 0 {
		log.Fatalf("App data file is invalid:\n%v", strings.Join(validationErrors, "\n"))
	}

	permissionRepository := permissionRepositoryConstructor.New(appData)
	defaultRolesVersionRepository := defaultRolesVersionRepositoryConstructor.New(appData)

	permissionService := permission.New(permissionRepository, defaultRolesVersionRepository)

	web.MinimalInit(
		r,

		permissionService,

		sharedMiddlewares...,
	)

	return *serverPort
}
