package routers

import (
	"net/http"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
)

//RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine, db orm.Ormer) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	//Add All route
	//TestRoutes(route, db)
}
