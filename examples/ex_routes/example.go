package routers

import (
	"gin-boilerplate/examples/ex_controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ExamplesRoutes(route *gin.Engine, db *gorm.DB) {
	ctrl := ex_controllers.Controller{DB: db}
	v1 := route.Group("/v1/examples")
	v1.GET("test/", ctrl.GetHasManyRelationData)

}
