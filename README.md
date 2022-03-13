# Go Boilerplate
An API boilerplate written in Golang with Gin Framework

### Docker Development with Live Reload
- run `make dev`

### Configuration Manage
- Manage from [config.yml](config.yml) file
```yaml
database:
  driver: "postgres"
  dbname: "test_pg_go"
  username: "mamun"
  password: "123"
  host: "postgres_db" # use "localhost" for local development, `postgres_db` for docker
  port: "5432"
  log_mode: true # SQL logger , false in production

server:
  host: "0.0.0.0"
  port: "8000"
  secret: "secret"
  environment: "dev" #debug logger ,use `prod` in production
  request:
    timeout: 100
```
### Local Setup Instruction
Follow these steps:
- Configuration manage from [config.yml](config.yml) file
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally run `go run main.go` or `go build main.go` and run `./main`
- The application should be available and running on [0.0.0.0:8000/health](http://0.0.0.0:8000/health)

### Container Build
- run `make build`
- 
### Build for Production
- run `make production`

### Under the hood
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps (Docker Development)
