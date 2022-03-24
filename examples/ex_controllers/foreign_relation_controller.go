package ex_controllers

import (
	"gin-boilerplate/examples/ex_models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetNormalData get normal data if added pagination see example_controller
func (base *Controller) GetNormalData(ctx *gin.Context) {
	var categories []ex_models.Category
	base.DB.Find(&categories)
	ctx.JSON(http.StatusOK, gin.H{"data": categories})

}

// GetForeignRelationData Get Foreign Data with Preload
func (base *Controller) GetForeignRelationData(ctx *gin.Context) {
	var articles []ex_models.Article
	base.DB.Preload("Category").Find(&articles)
	ctx.JSON(http.StatusOK, &articles)

}
