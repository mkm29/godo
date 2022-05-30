package main

import (
	"github.com/mkm29/godo/api/v1/router"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	log.Info("Starting Todolist API server")
	router.ServeRouter()
}
