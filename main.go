package main

import (
	"gin-boilerplate/pkg/config"
	"gin-boilerplate/pkg/database"
	"gin-boilerplate/pkg/logger"
	"gin-boilerplate/routers"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	if err := database.SetupConnection(); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	router := routers.SetupRoute()

	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
