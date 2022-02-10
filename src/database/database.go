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

var DB = *dbConnection()

// dbConnection create database connection
func dbConnection() *bun.DB {
	// connection string and open database
	dns := dbConfig.Database()
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dns)))
	db := bun.NewDB(sqlDb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false),
		bundebug.FromEnv("BUN_DEBUG")))
	log.Println("Database Successfully connected!")
	return db
}

func _Save(ctx context.Context, model interface{}) error {
	_, err := DB.NewInsert().Model(model).Exec(ctx)
	return err

}

func _Get(ctx context.Context, model interface{}, orderBy string, limit int, offSet int) error {
	query := DB.NewSelect().Column().Model(model).OrderExpr(orderBy).Limit(limit).Offset(offSet).Scan(ctx)
	return query
}
