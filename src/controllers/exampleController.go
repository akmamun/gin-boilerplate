package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pkg/src/logger"
	"pkg/src/models"
)

func (base *Controller) CreateExample(ctx *gin.Context) {
	example := new(models.Example)

	err := ctx.ShouldBindJSON(&example)
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}
	err = base.DB.Create(&example).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.AbortWithStatus(404)
		return
	}
	ctx.JSON(200, &example)
}

func (base *Controller) GetExamples(ctx *gin.Context) {
	var example []models.Example
	err := base.DB.Find(&example).Error

	if err != nil {
		logger.Errorf("error: %v", err)
	}

	ctx.JSON(http.StatusOK, example)

}
