package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mini-roles-backend/source/domains/files/services/validation"
	"mini-roles-backend/source/domains/permission/repositories/permission_memory"
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

func minimalInit(r *gin.Engine) int {
	appDataBytes, err := ioutil.ReadFile(*appDataPath)
	if err != nil {
		log.Fatalf("Unable to read app data file by path %s: %v", *appDataPath, err)
	}

	appData, validationErrors := validation.Validate(appDataBytes)
	if len(validationErrors) > 0 {
		log.Println("App data file is invalid:")
		log.Fatalln(strings.Join(validationErrors, "\n"))
	}

	permissionRepository := permission_memory.New(appData)
	permissionService := permission.New(permissionRepository)

	web.MinimalInit(
		r,

		permissionService,
	)

	return *serverPort
}
