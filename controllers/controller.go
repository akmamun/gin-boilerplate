package controllers

import (
	"gorm.io/gorm"
)

type Controller struct {
	//rep repository.Repository
	DB *gorm.DB
}
