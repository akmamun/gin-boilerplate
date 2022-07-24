package main

import (
	"gin-boilerplate/migrations"
	"gin-boilerplate/pkg/config"
	"gin-boilerplate/pkg/database"
	"gin-boilerplate/pkg/logger"
	"gin-boilerplate/routers"
	"github.com/spf13/viper"
	"time"
)

func main() {

	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	if err := database.DbConnection(); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}
	//TODO changes in later
	migrations.Migrate()
	router := routers.SetupRoute()

	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
