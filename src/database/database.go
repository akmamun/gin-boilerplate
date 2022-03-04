package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"pkg/src/models"
)

var (
	DB    *gorm.DB
	err   error
	DBErr error
)

type Database struct {
	*gorm.DB
}

func configuration() string {
	dbname := viper.GetString("database.dbname")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	sslMode := viper.GetString("database.sslmode")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, username, password, dbname, port, sslMode,
	)
	return dsn
}

// Connection create database connection
func Connection() error {
	var db = DB
	dsn := configuration()

	logMode := viper.GetBool("database.logmode")
	loglevel := logger.Silent

	if logMode {
		loglevel = logger.Info
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})

	if err != nil {
		DBErr = err
		log.Println("DbConfiguration connection error")
		return err
	}

	err = db.AutoMigrate(&models.Example{})
	if err != nil {
		return err
	}
	DB = db

	return nil

}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}

// GetDBErr helps you to get a connection
func GetDBErr() error {
	return DBErr
}

//func _GetSelected() {
//	DB.Select("Name", "Age", "CreatedAt").Create(&user)
//
//}
//func _update(model interface{}, fieldData schema.Field) {
//	data := (&model, fieldData)
//	data.field = model
//}
