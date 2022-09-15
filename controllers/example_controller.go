package controllers

import (
	"gin-boilerplate/models"
	"gin-boilerplate/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetData(ctx *gin.Context) {
	var example []*models.Example
	repository.Get(&example)
	ctx.JSON(http.StatusOK, &example)

}
