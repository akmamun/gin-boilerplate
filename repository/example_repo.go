package repository

//
//import (
//	"github.com/gin-gonic/gin"
//	"gorm.io/gorm"
//	"net/http"
//	"pkg/src/helpers/pagination"
//	"pkg/src/ex_models"
//	"strconv"
//)
//
//
//func (db *gorm.DB) GetExamples(ctx *gin.Context) {
//	var example []ex_models.Example
//
//	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
//	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "0"))
//	query := ctx.DefaultQuery("query", "")
//
//	paginateData := pagination.Pagination(&pagination.Param{
//		DB:    base.DB,
//		Page:  int64(page),
//		Limit: int64(limit),
//		//OrderBy: "id desc",
//		Search: query,
//	}, &example)
//
//	ctx.JSON(http.StatusOK, paginateData)
//
//}
