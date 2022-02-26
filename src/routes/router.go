package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pkg/src/controllers"
)

func Routers(db *gorm.DB) *gin.Engine {
	route := gin.New()
	route.Use(gin.Logger())
	api := controllers.Controller{DB: db}

	//Todo group route

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": nil, "message": "system is live "})
	})
	route.POST("/example/", api.CreateExample)
	route.GET("/example/", api.GetExample)

	return route
}
