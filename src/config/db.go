package config

import (
	"fmt"
	"strconv"

	env "go-pg/src/helpers"
)

// database configrations and return connection string
func Database() string {
	host := env.GetEnv("DB_HOST", "localhost")
	port, _ := strconv.Atoi(env.GetEnv("DB_PORT", "5432")) //parse int
	user := env.GetEnv("DB_USER", "postgres")
	dbname := env.GetEnv("DB_NAME", "go_pg")
	password := env.GetEnv("DB_PASSWORD", "")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return psqlconn
}
