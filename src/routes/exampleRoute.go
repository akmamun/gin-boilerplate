package routes

import (
	"github.com/gin-gonic/gin"
	controller "pkg/src/controllers"
)

func ExampleRoutes(route *gin.Engine) {
	route.GET("/example/", controller.ExampleResponse)
}
