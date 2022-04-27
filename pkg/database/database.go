package database

import (
	"gin-boilerplate/pkg/config"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/lib/pq"
)

var (
	DB    orm.Ormer
	err   error
	DBErr error
)

type Database struct {
	orm.Ormer
}

// Connection create database connection
func Connection() error {
	var db = DB
	
	l := logs.GetLogger()
	l.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm")
	
    logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")
	
	// db, err = orm.Open(postgres.Open(dsn), &gorm.Config{
		// 	Logger: logger.Default.LogMode(loglevel),
		// })

	// if err != nil {
	// 	DBErr = err
	// 	log.Println("Db connection error")
	// 	return err
	// }
	err = orm.RunSyncdb("default", true, true)

	// err = db.AutoMigrate(migrationModels...)

	if err != nil {
		return err
	}
	DB = db

	return nil

}

func init() {
	dsn := config.DbConfiguration()
	orm.RegisterDriver("postgres", orm.DRPostgres)

	orm.RegisterDataBase("default", "postgres", dsn)
}

// GetDB connection
func GetDB() (orm.Ormer) {
	// db,_:= orm.GetDB()
	return DB
}

// GetDBError connection error
func GetDBError() error {
	return DBErr
}
