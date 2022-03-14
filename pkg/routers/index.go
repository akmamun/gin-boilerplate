package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

//RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine, db *gorm.DB) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })

	//Add All route
	TestRoutes(route, db)
}
