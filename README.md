# Gin Gorm Boilerplate
An API boilerplate written in Golang with Gin Framework


### Docker Development with Live Reload
- run `make dev`

### Local Setup Instruction
Follow these steps:
- Configuration manage from [config.yml](config.yml) file
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally run `go run main.go` or `go build main.go` and run `./main`
- The application should be available and running on [0.0.0.0:8000/health](http://0.0.0.0:8000/health)

### Build for Production
- run `make production`

### Under the hood
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps (Docker Development)
