package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine, db *gorm.DB) {
	TestRoutes(route, db)
}
