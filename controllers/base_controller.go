package controllers

import (
	"gorm.io/gorm"
)

type BaseController struct {
	//rep repository.Repository
	DB *gorm.DB
}
