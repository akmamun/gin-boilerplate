package database

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	dbConfig "go-pg/src/config"
	"log"
)

var Db = *dbConnection()

// dbConnection create database connection
func dbConnection() *bun.DB {
	// connection string and open database
	dns := dbConfig.Database()
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dns)))
	db := bun.NewDB(sqlDb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	log.Println("Database Successfully connected!")
	return db
}

func Save(ctx context.Context, model interface{}) error {
	_, err := Db.NewInsert().Model(model).Exec(ctx)
	return err

}

func Get(ctx context.Context, model interface{}) error {
	query := Db.NewSelect().Column().Model(model).Scan(ctx)
	return query
}
