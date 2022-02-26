package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	response "pkg/src/applibs"
	"pkg/src/database"
	"pkg/src/models"
)

func (base *Controller) CreateExample(ctx *gin.Context) {
	example := new(models.Example)

	err := ctx.ShouldBindJSON(&example)
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	saveData, err := database.Save(base.DB, &example)
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ctx.JSON(200, saveData)
}

func (base *Controller) GetExample(ctx *gin.Context) {
	var example []models.Example

	//pagination := response.LimitOffsetOrPagination(ctx)
	//fmt.Print(pagination)
	database.GetAll(base.DB, &example)
	ctx.JSON(http.StatusOK, response.SuccessfullyGet(&example))

}

//
//func ExampleResponse(ctx *gin.Context) {
//	var example models.Example
//
//	if err := ctx.ShouldBindJSON(&example); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	//example = models.Example{Data: example.Data}
//	database.Save(&example)
//	ctx.JSON(http.StatusCreated, response.SuccessfullyCreated())
//
//}
//
