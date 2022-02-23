package controllers

import (
	"github.com/gin-gonic/gin"
	helpers "pkg/src/helpers"
)

func ExampleResponse(ctx *gin.Context) {
	data := "test"
	helpers.SuccessResponse(ctx, data)
	return
}
