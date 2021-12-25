package config

import (
	"fmt"
	env "go-pg/src/helpers"
	"strconv"
)

// Database configurations and return connection string
func Database() string {
	host := env.GetEnv("DB_HOST", "localhost")
	port, _ := strconv.Atoi(env.GetEnv("DB_PORT", "5432")) //parse int
	user := env.GetEnv("DB_USER", "")
	dbname := env.GetEnv("DB_NAME", "")
	password := env.GetEnv("DB_PASSWORD", "")

	dbConn := fmt.Sprintf("host=%s:port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return dbConn
}
