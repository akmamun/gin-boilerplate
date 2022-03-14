package controllers

import (
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/helpers/pagination"
	"gin-boilerplate/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (base *Controller) CreateExample(ctx *gin.Context) {
	example := new(models.Example)

	err := ctx.ShouldBindJSON(&example)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = base.DB.Create(&example).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &example)
}

func (base *Controller) GetExamples(ctx *gin.Context) {
	var example []models.Example

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "0"))
	query := ctx.DefaultQuery("query", "")

	//db := base.DB.Where("")
	paginateData := pagination.Pagination(&pagination.Param{
		DB:    base.DB, //db
		Page:  int64(page),
		Limit: int64(limit),
		//OrderBy: "id desc",
		Search: query,
	}, &example)

	ctx.JSON(http.StatusOK, paginateData)

}
