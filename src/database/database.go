package database

import (
	"fmt"
	"log"
	dbConfig "pkg/src/config"
	"pkg/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB = *DbConnection()

// dbConnection create database connection
func DbConnection() *gorm.DB {
	// connection string and open database
	dns := dbConfig.Database()
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dns,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
	if err != nil {
		log.Println("Database connection error")
		log.Fatal()

	}
	//db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Example{})

	return db
}

func Save(model interface{}) interface{} {
	query := DB.Create(model)
	if query.Error != nil {
		fmt.Println(query.Error)
	}
	return query
}

func _GetOne(model interface{}, field string, value interface{}) interface{} {
	query := DB.Where(fmt.Sprintf("%v = ?", field), value).Find(&model)
	if query.Error != nil {
		fmt.Println(query.Error)
	}
	return query
}

func _GetFirst(model interface{}, field string) interface{} {
	query := DB.First(&model, field)
	if query.Error != nil {
		fmt.Println(query.Error)
	}
	return query
}

func _GetLast(model interface{}, fieldData schema.Field) interface{} {
	query := DB.Last(&model, fieldData)
	if query.Error != nil {
		fmt.Println(query.Error)
	}
	return query
}
func GetAll(model interface{}, limit, offSet int) interface{} {
	query := DB.Limit(limit).Offset(offSet).Find(model)
	if query.Error != nil {
		fmt.Println(query.Error)
	}
	return query
}

//func _GetSelected() {
//	DB.Select("Name", "Age", "CreatedAt").Create(&user)
//
//}
//func _update(model interface{}, fieldData schema.Field) {
//	data := (&model, fieldData)
//	data.field = model
//}
