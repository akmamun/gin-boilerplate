package config

import (
	"fmt"
	env "pkg/src/helpers"
)

// Database configurations and return connection string
func Database() string {

	user := env.GetEnv("DB_USER", "postgres")
	host := env.GetEnv("DB_HOST", "localhost")
	port := env.GetEnv("DB_PORT", "5432")
	dbname := env.GetEnv("DB_NAME", "")
	password := env.GetEnv("DB_PASSWORD", "")

	psqlConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	return psqlConn
}
