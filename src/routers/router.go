package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
	"pkg/src/routers/middleware"
)

func Routes(db *gorm.DB) *gin.Engine {

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

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "active": true})
	})

	return router
}
