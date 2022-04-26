package config

import (
	"fmt"
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

	appServer := fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))
	return appServer
}
