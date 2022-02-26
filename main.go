package main

import (
	"fmt"
	"github.com/spf13/viper"
	"pkg/src/config"
	"pkg/src/database"
	"pkg/src/logger"
	"pkg/src/routes"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config.ConfigSetup() error: %s", err)
	}

	if err := database.DbConnection(); err != nil {
		logger.Fatalf("database.DbConnection error: %s", err)
	}

	db := database.GetDB()
	//
	route := routes.Routers(db)

	host := "127.0.0.1"
	if h := viper.GetString("server.host"); h != "" {
		host = h
	}
	logger.Infof("Server is starting at %s:%s", host, viper.GetString("server.port"))
	fmt.Printf("Server is starting at %s:%s.", host, viper.GetString("server.port"))
	logger.Fatalf("%v", route.Run(host+":"+viper.GetString("server.port")))

}
