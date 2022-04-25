package config

import (
	"github.com/spf13/viper"
)

type ServerConfiguration struct {
	Port                 string
	Secret               string
	LimitCountPerRequest int64
}

func ServerConfig() string {
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", "8000")

	host := viper.GetString("server.host")
	port := viper.GetString("server.port")

	return host + ":" + port

}
