package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	response "pkg/src/applibs"
	"pkg/src/database"
	"pkg/src/models"
)

func ExampleResponse(ctx *gin.Context) {
	var example models.Example

	if err := ctx.ShouldBindJSON(&example); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//example = models.Example{Data: example.Data}
	database.Save(&example)
	ctx.JSON(http.StatusCreated, response.SuccessfullyCreated())

}

func ExampleGetResponse(ctx *gin.Context) {
	var example []models.Example
	database.GetAll(&example, 1, 10)
	ctx.JSON(http.StatusOK, response.SuccessfullyGet(&example))

}
