package database

import (
	"gin-boilerplate/pkg/config"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	DB    *gorm.DB
	err   error
	DBErr error
)

type Database struct {
	*gorm.DB
}

// Connection create database connection
func Connection() error {
	var db = DB
	dsn := config.DbConfiguration()

	logMode := viper.GetBool("database.log_mode")
	loglevel := logger.Silent
	if logMode {
		loglevel = logger.Info
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})

	if err != nil {
		DBErr = err
		log.Println("Db connection error")
		return err
	}

	err = db.AutoMigrate(migrationModels...)

	if err != nil {
		return err
	}
	DB = db

	return nil

}

// GetDB connection
func GetDB() *gorm.DB {
	return DB
}

// GetDBError connection error
func GetDBError() error {
	return DBErr
}
