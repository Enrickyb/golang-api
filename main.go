package main

import (
	"github.com/Enrickyb/golang-api/config"
	"github.com/Enrickyb/golang-api/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")
	//Initialize configs

	err := config.Init()
	if err != nil {
		logger.ErrorF("config inicialization error: %v", err)
		return
	}

	router.Initialize()

}
