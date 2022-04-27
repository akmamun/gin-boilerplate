package config

import (
	"fmt"
)


type DatabaseConfiguration struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

func DbConfiguration() string {
	// dbname := viper.GetString("database.dbname")
	// username := viper.GetString("database.username")
	// password := viper.GetString("database.password")
	// host := viper.GetString("database.host")
	// port := viper.GetString("database.port")
	// sslMode := viper.GetString("database.sslmode")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"192.168.10.160", "mamun", "123", "test_pg_go", "5432", "disable",
	)
	fmt.Println("######################################################")
	fmt.Println(dsn)
	// dns := "mamun:123@192.168.10.160:5432/test_pg_go?charset=utf8"
	return dsn
}
