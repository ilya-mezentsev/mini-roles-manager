package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var mode = flag.String("mode", "full", "Set application mode (full or minimal)")

func init() {
	flag.Parse()
	configureLogging()
}

func main() {
	r := gin.New()
	middlewares := []gin.HandlerFunc{
		gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: logFormatter,
			Output:    logWriter{},
		}),

		gin.Recovery(),
	}

	var port int
	switch *mode {
	case "full":
		port = fullInit(r, middlewares...)

	case "minimal":
		port = minimalInit(r, middlewares...)

	default:
		log.Fatalf("Unknown mode: %s", *mode)
	}

	log.Infof("Starting in mode: %s", *mode)
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}
}
