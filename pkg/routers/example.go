package routers

import (
	"gin-boilerplate/controllers"
	"net/http"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
)

func TestController(ctx *gin.Context) {
	ctx.String(http.StatusOK, "test data ok")
}

func TestRoutes(route *gin.Engine, db orm.Ormer) {
	ctrl := controllers.Controller{DB: db}
	v1 := route.Group("/v1")
	v1.POST("/example/", ctrl.CreateExample)
	// v1.GET("/example/", ctrl.GetExamples)
	// v1.GET("test/", TestController)

}
