package controllers

import (
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/helpers/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (base *Controller) GetPaginatedData(ctx *gin.Context) {
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
