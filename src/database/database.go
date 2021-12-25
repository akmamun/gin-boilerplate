package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"go-pg/src/config"
)

func dbConnection() (db *bun.DB) {
	// connection string and open database
	dsn := config.Database()
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	fmt.Println("Successfully connected!")
	return db

}

func Save(ctx context.Context, data interface{}) interface{} {
	db := dbConnection()
	res, err := db.NewInsert().Model(data).Exec(ctx)
	if err != nil {
		panic(err)
	}
	return res
}
