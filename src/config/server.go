package config

type ServerConfiguration struct {
	Port                 string
	Secret               string
	LimitCountPerRequest float64
}
