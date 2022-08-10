package controllers

import (
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/helpers"
	"encoding/json"
	"net/http"
)

func GetData(w http.ResponseWriter, request *http.Request) {
	var example []models.Example
	models.GetAll(&example)
	helpers.SuccessResponse(w, &example)
}