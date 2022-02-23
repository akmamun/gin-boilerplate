package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	dbConfig "pkg/src/config"
)

//var DB = *dbConnection()

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
	}
	log.Println("Database Successfully connected!")
	//db.AutoMigrate(&models.Book{})
	return db
}

//
//func _Save(ctx context.Context, model interface{}) error {
//	_, err := DB.NewInsert().Model(model).Exec(ctx)
//	return err
//
//}
//
//func _Get(ctx context.Context, model interface{}, orderBy string) error {
//	query := DB.NewSelect().Column().Model(model).OrderExpr(orderBy).Scan(ctx)
//	return query
//}
