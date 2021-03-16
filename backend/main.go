package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
)

var (
	configFilePath = flag.String("config", "/dev/null", "Set path to configs file")
)

func init() {
	flag.Parse()
}

func main() {
	log.Infof("Got config path: %s", *configFilePath)
}
