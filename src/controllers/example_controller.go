package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pkg/src/helpers"
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

//func (base *Controller) GetExamples(ctx *gin.Context) {
//	var example []models.Example
//	err := base.DB.Find(&example).Error
//
//	if err != nil {
//		logger.Errorf("error: %v", err)
//	}
//
//	ctx.JSON(http.StatusOK, example)
//
//}

func (base *Controller) GetExamples(ctx *gin.Context) {
	var example []models.Example
	err := base.DB.Find(&example).Error
	//err := base.rep.Find(&example)
	if err != nil {
		logger.Errorf("error: %v", err)
	}
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "0"))
	p, _ := (&helpers.Param{
		Page:  page,
		Limit: limit,
		Path:  ctx.FullPath(),
		Sort:  "id asc",
	}).Paginate(base.DB, &example)
	ctx.JSON(http.StatusOK, p)

}
