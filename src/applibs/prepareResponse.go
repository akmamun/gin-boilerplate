package applibs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessfullyCreated() interface{} {
	response := make(map[string]interface{})
	response["success"] = true
	response["message"] = "successfully created"
	response["data"] = nil
	return response
}

func SuccessfullyGet(data interface{}) interface{} {
	response := make(map[string]interface{})
	response["success"] = true
	response["message"] = "successfully return"
	response["data"] = data
	return response
}

func ErrorResponse(ctx *gin.Context, data interface{}) interface{} {
	response := make(map[string]interface{})
	response["message"] = "Data validation error"
	response["data"] = data
	ctx.JSON(http.StatusBadRequest, response)
	return response
}