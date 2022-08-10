package repository

import (
	"gin-boilerplate/pkg/database"
	"gin-boilerplate/pkg/logger"
)

func Save(value interface{}) interface{} {
	err := database.GetDB().Create(value).Error
	if err != nil {
		logger.Errorf("error, not save data %v", err)
	}
	return err
}

func Get(value interface{}) interface{} {
	err := database.GetDB().Find(value).Error
	return err
}

func GetOne(value interface{}) interface{} {
	err := database.GetDB().Last(value).Error
	return err
}

func Update(value interface{}) interface{} {
	err := database.GetDB().Find(value).Error
	return err
}
