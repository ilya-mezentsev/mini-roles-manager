package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type logWriter struct {
}

func (l logWriter) Write(p []byte) (n int, err error) {
	log.Info(string(p))

	return len(p), nil
}

func logFormatter(params gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"%s|%15s|%15d %s \"%s\"",
		params.ClientIP,
		params.Latency,
		params.StatusCode,
		params.Method,
		params.Path,
	)
}

func configureLogging() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
