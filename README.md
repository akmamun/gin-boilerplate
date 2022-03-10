# Gin Gorm Boilerplate
An API boilerplate written in Golang with Gin Framework

### Docker Development with Live Reload
- Run `make dev`

### Local Installation
- configuration manage from [config.yml](config.yml) file
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally Run`go run main.go` or `go build main.go` and run `./main`
- The application should be available and running on 0.0.0.0:8000

### Build for Production
- Run`make production`

### Under the hood
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
