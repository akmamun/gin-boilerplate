package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	controller "pkg/src/controllers"
)

func Routers(db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	api := controller.Controller{DB: db}

	//Todo group route

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": nil, "message": "system is live "})
	})
	router.POST("/example/", api.CreatePost)
	//route.POST("/example/", controller.ExampleResponse)

	return router
}
