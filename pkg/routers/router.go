package routers

import (
	"gin-boilerplate/pkg/routers/middleware"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Routes(db orm.Ormer) *gin.Engine {

	environment := viper.Get("server.environment")
	if environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	RegisterRoutes(router, db) //routes register

	return router
}
