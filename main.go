package main

import (
	"github.com/spf13/viper"
	"pkg/src/config"
	"pkg/src/database"
	"pkg/src/logger"
	"pkg/src/routers"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config.SetupConfig() error: %s", err)
	}

	if err := database.Connection(); err != nil {
		logger.Fatalf("database.DbConnection error: %s", err)
	}

	db := database.GetDB()
	router := routers.Routes(db)

	host := viper.GetString("server.host")
	if host == "" {
		host = "0.0.0.0"
	}
	logger.Infof("Server is starting at %s:%s", host, viper.GetString("server.port"))
	logger.Fatalf("%v", router.Run(host+":"+viper.GetString("server.port")))

}
