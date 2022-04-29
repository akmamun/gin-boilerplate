package config

import (
	"fmt"
	"github.com/spf13/viper"
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

func DbConfiguration() (string, string) {
	masterDBName := viper.GetString("MASTER_DB_NAME")
	masterDBUser := viper.GetString("MASTER_DB_USER")
	masterDBPassword := viper.GetString("MASTER_DB_PASSWORD")
	masterDBHost := viper.GetString("MASTER_DB_HOST")
	masterDBPort := viper.GetString("MASTER_DB_PORT")
	masterDBSslMode := viper.GetString("MASTER_SSL_MODE")

	replicaDBName := viper.GetString("REPLICA_DB_NAME")
	replicaDBUser := viper.GetString("REPLICA_DB_USER")
	replicaDBPassword := viper.GetString("REPLICA_DB_PASSWORD")
	replicaDBHost := viper.GetString("REPLICA_DB_HOST")
	replicaDBPort := viper.GetString("REPLICA_DB_PORT")
	replicaDBSslMode := viper.GetString("REPLICA_SSL_MODE")

	masterDBDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		masterDBHost, masterDBUser, masterDBPassword, masterDBName, masterDBPort, masterDBSslMode,
	)

	replicaDBDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		replicaDBHost, replicaDBUser, replicaDBPassword, replicaDBName, replicaDBPort, replicaDBSslMode,
	)
	return masterDBDSN, replicaDBDSN
}
