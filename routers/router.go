package routers

import (
	"gin-boilerplate/routers/middleware"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Routes(db orm.Ormer) *gin.Engine {

	environment := viper.GetBool("DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	allowedHosts := viper.GetString("ALLOWED_HOSTS")
	router := gin.New()
	router.SetTrustedProxies([]string{allowedHosts})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	RegisterRoutes(router, db) //routes register

	return router
}
