package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResponse(ctx *gin.Context, data interface{}) interface{} {
	response := make(map[string]interface{})
	response["message"] = "successfully return"
	response["data"] = data
	ctx.JSON(http.StatusOK, response)
	return response
}

func ErrorResponse(ctx *gin.Context, data interface{}) interface{} {
	response := make(map[string]interface{})
	response["message"] = "Data validation error"
	response["data"] = data
	ctx.JSON(http.StatusBadRequest, response)
	return response
}
