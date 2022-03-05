package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pkg/src/helpers/pagination"
	"pkg/src/logger"
	"pkg/src/models"
	"strconv"
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
