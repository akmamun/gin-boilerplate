package routers

import (
	"gin-boilerplate/routers/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) *gin.Engine {

	environment := viper.Get("server.environment")
	if environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	allowHosts := viper.GetString("server.allow_hosts")

	router := gin.New()
	router.SetTrustedProxies([]string{allowHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	RegisterRoutes(router, db) //routes register

	return router
}
