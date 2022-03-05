package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(route *gin.Engine, db *gorm.DB) {
	TestRoutes(route, db)
}
