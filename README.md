# Go Boilerplate
An API boilerplate written in Golang with Gin Framework and Gorm 
### Motivation
Write restful API with fast development and developer friendly

## Table of Contents
- [Motivation](#motivation)
- [Configuration Manage](#configuration-manage)
  - [ENV Manage](#env-manage)
  - [Server Configuration](#server-configuration)
  - [Database Configuration](#database-configuration)
  - [PgAdmin](#pg-admin)
- [Installation](#installation)
  - [Local Setup Instruction](#local-setup-instruction)
  - [Develop Application in Docker with Live Reload](#develop-application-in-docker-with-live-reload)
- [Middlewares](#middlewares)
- [Boilerplate Structure](#boilerplate-structure)
- [Code Examples](#examples)
- [Let's Build an API](#lets-build-an-api)
- [Deployment](#deployment)
  - [Container Development Build](#container-development-build)
  - [Container Production Build and Up](#container-production-build-and-up)
- [Useful Commands](#useful-commands)
- [ENV YAML Configure](#env-yaml-configure)
- [Use Packages](#use-packages)

### Configuration Manage
#### ENV Manage

- Default ENV Configuration Manage from `.env`. sample file `.env.example`
```text
# Server Configuration
SECRET=h9wt*pasj6796j##w(w8=xaje8tpi6h*r&hzgrz065u&ed+k2)
DEBUG=True # `False` in Production
ALLOWED_HOSTS=0.0.0.0
SERVER_HOST=0.0.0.0
SERVER_PORT=8000

# Database Configuration
MASTER_DB_NAME=test_pg_go
MASTER_DB_USER=mamun
MASTER_DB_PASSWORD=123
MASTER_DB_HOST=postgres_db
MASTER_DB_PORT=5432
MASTER_DB_LOG_MODE=True # `False` in Production
MASTER_SSL_MODE=disable

REPLICA_DB_NAME=test_pg_go
REPLICA_DB_USER=mamun
REPLICA_DB_PASSWORD=123
REPLICA_DB_HOST=localhost
REPLICA_DB_PORT=5432
REPLICA_DB_LOG_MODE=True # `False` in Production
REPLICA_SSL_MODE=disable
```
- Server `DEBUG` set `False` in Production
- Database Logger `MASTER_DB_LOG_MODE` and `REPLICA_DB_LOG_MODE`  set `False` in production
- If ENV Manage from YAML file add a config.yml file and configuration [db.go](config/db.go) and [server.go](config/server.go). See More [ENV YAML Configure](#env-yaml-configure)

#### Server Configuration
- Use [Gin](https://github.com/gin-gonic/gin) Web Framework

#### Database Configuration
- Use [GORM](https://github.com/go-gorm/gorm) as an ORM
- Use database `MASTER_DB_HOST` value set as `localhost` for local development, and use `postgres_db` for docker development 
#### PG Admin
- Check  PG Admin on [http://0.0.0.0:5050/browser/](http://0.0.0.0:5050/browser/)
- Login with Credential Email `admin@admin.com` Password `root`
- Connect Database Host as `postgres_db`, DB Username and Password as per `.env` set
- Note: if not configure `.env`, default Username `mamun` and password `123`

### Installation
#### Local Setup Instruction
Follow these steps:
- Copy [.env.example](.env.example) as `.env` and configure necessary values
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally run `go run main.go` or `go build main.go` and run `./main`
- Check Application health available on [0.0.0.0:8000/health](http://0.0.0.0:8000/health)

#### Develop Application in Docker with Live Reload
Follow these steps:
- Make sure install the latest version of docker and docker-compose
- Docker Installation for your desire OS https://docs.docker.com/engine/install/ubuntu/
- Docker Composer Installation https://docs.docker.com/compose/install/
- Run and Develop `make dev`
- Check Application health available on [0.0.0.0:8000/health](http://0.0.0.0:8000/health)

### Middlewares
- Use Gin CORSMiddleware
```go
router := gin.New()
router.Use(gin.Logger())
router.Use(gin.Recovery())
router.Use(middleware.CORSMiddleware())
```

### Boilerplate Structure
<pre>├── <font color="#3465A4"><b>config</b></font>
│   ├── config.go
│   ├── db.go
│   └── server.go
├── <font color="#3465A4"><b>controllers</b></font>
│   └── example_controller.go
├── docker-compose-dev.yml
├── docker-compose-prod.yml
├── Dockerfile
├── Dockerfile-dev
├── go.mod
├── go.sum
├── <font color="#3465A4"><b>helpers</b></font>
│   ├── response.go
│   └── search.go
├── <font color="#3465A4"><b>infra</b></font>
│   ├── <font color="#3465A4"><b>database</b></font>
│   │   └── database.go
│   └── <font color="#3465A4"><b>logger</b></font>
│       └── logger.go
├── LICENSE
├── main.go
├── Makefile
├── <font color="#3465A4"><b>migrations</b></font>
│   └── migration.go
├── <font color="#3465A4"><b>models</b></font>
│   └── example_model.go
├── README.md
├── <font color="#3465A4"><b>repository</b></font>
│   └── sql_repo.go
├── <font color="#3465A4"><b>routers</b></font>
│   ├── index.go
│   ├── <font color="#3465A4"><b>middleware</b></font>
│   │   └── cors.go
│   └── router.go
</pre>

### Examples
- More Example [gin-boilerplate-examples](https://github.com/akmamun/gin-boilerplate-examples)

### Let's Build an API

1. [models](models) folder add a new file name `example_model.go`

```go
package models

import (
	"time"
)

type Example struct {
	Id        int        `json:"id"`
	Data      string     `json:"data" binding:"required"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}
// TableName is Database Table Name of this model
func (e *Example) TableName() string {
	return "examples"
}
```
2. Add Model to [migration](pkg/database/migration.go)
```go
package migrations

import (
	"gin-boilerplate/infra/database"
	"gin-boilerplate/models"
)

// Migrate Add list of model add for migrations
func Migrate() {
	var migrationModels = []interface{}{&models.Example{}}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}

```
3. [controller](controllers) folder add a file `example_controller.go`
- Create API Endpoint 
- Write Database Operation in Repository and use them from controller
```go
package controllers

import (
  "gin-boilerplate/models"
  "gin-boilerplate/repository"
  "github.com/gin-gonic/gin"
  "net/http"
)

func GetData(ctx *gin.Context) {
  var example []*models.Example
  repository.Get(&example)
  ctx.JSON(http.StatusOK, &example)

}
func Create(ctx *gin.Context) {
  example := new(models.Example)
  repository.Save(&example)
  ctx.JSON(http.StatusOK, &example)
}
```
4. [routers](routers) folder add a file `example.go`
```go
package routers

import (
  "gin-boilerplate/controllers"
  "github.com/gin-gonic/gin"
  "net/http"
)

func RegisterRoutes(route *gin.Engine) {
  route.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"live": "ok"}) })
  //added new
  route.GET("/v1/example/", controllers.GetData)
  route.POST("/v1/example/", controllers.Create)

  //Add All route
  //TestRoutes(route)
}
```
- Congratulation, your new endpoint `0.0.0.0:8000/v1/example/`

### Deployment
#### Container Development Build
- Run `make build`

#### Container Production Build and Up
- Run `make production`

#### ENV Yaml Configure
```yaml
database:
  driver: "postgres"
  dbname: "test_pg_go"
  username: "mamun"
  password: "123"
  host: "postgres_db" # use `localhost` for local development
  port: "5432"
  ssl_mode: disable
  log_mode: false

server:
  host: "0.0.0.0"
  port: "8000"
  secret: "secret"
  allow_hosts: "localhost"
  debug: false #use `false` in production
  request:
    timeout: 100
```
- [Server Config](config/server.go)
```go
func ServerConfig() string {
viper.SetDefault("server.host", "0.0.0.0")
viper.SetDefault("server.port", "8000")
appServer := fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))
return appServer
}
```
- [DB Config](config/db.go)
```go
func DbConfiguration() string {
	
dbname := viper.GetString("database.dbname")
username := viper.GetString("database.username")
password := viper.GetString("database.password")
host := viper.GetString("database.host")
port := viper.GetString("database.port")
sslMode := viper.GetString("database.ssl_mode")

dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
host, username, password, dbname, port, sslMode)
return dsn
}
```

### Useful Commands

- `make dev`: make dev for development work
- `make build`: make build container
- `make production`: docker production build and up
- `clean`: clean for all clear docker images

### Use Packages
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps (Docker Development)

