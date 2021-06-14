package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var mode = flag.String("mode", "full", "Set application mode (default or minimal)")

func init() {
	flag.Parse()
}

func main() {
	r := gin.Default()

	var port int
	switch *mode {
	case "full":
		port = fullInit(r)

	case "minimal":
		port = minimalInit(r)

	default:
		log.Fatalf("Unknown mode: %s", *mode)
	}

	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}
}
