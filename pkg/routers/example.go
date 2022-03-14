package routers

import (
	"gin-boilerplate/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func TestController(ctx *gin.Context) {
	ctx.String(http.StatusOK, "test data ok")
}

func TestRoutes(route *gin.Engine, db *gorm.DB) {
	ctrl := controllers.Controller{DB: db}
	v1 := route.Group("/v1")
	v1.POST("/example/", ctrl.CreateExample)
	v1.GET("/example/", ctrl.GetExamples)
	v1.GET("test/", TestController)

}
