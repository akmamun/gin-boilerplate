# Go Boilerplate
An API boilerplate written in Golang with Gin Framework

## Table of Contents
- [Motivation](#motivation)
- [Configuration Manage](#configuration-manage)
  - [Server Configuration](#server-configuration)
  - [Database Configuration](#database-configuration)
- [Develop Application in Docker Compose with Live Reload](#develop-application-in-docker-compose-with-live-reload)
- [Local Setup Instruction](#local-setup-instruction)
- [Routes](#routes)
- [Middlewares](#middlewares)
- [Logging](#logging)
- [Boilerplate Structure](#boilerplate-structure)
- [Let's Build an API](#lets-build-an-api)
- [Code Examples](#code-examples)
- [Useful Commands](#useful-commands)
- [Container Development Build](#container-development-build)
- [Container Production Build and Up](#container-production-build-and-up)
- [Use Packages](#use-packages)

### Motivation
Write restful API with fast development and developer friendly

### Configuration Manage
- Manage from [config.yml](config.yml) file
- Use [Gin](https://github.com/gin-gonic/gin) Web Framework
- Server `environment` is Gin debug logger, use `prod` in production and `dev` in development mode

#### Server Configuration
```yaml
server:
  host: "0.0.0.0"
  port: "8000"
  secret: "secret"
  environment: "dev" #debug logger ,use `prod` in production
  request:
    timeout: 100
```

#### Database Configuration
- Use [GORM](https://github.com/go-gorm/gorm) as an ORM. you just need to configure config.yml file according to your setup.
- Use database `host` as `localhost` for local development, if docker use `postgres_db`
- Database `log_mode` is SQL logger, `false` in production and `true` in development mode
```yaml
database:
  driver: "postgres"
  dbname: "test_pg_go"
  username: "mamun"
  password: "123"
  host: "postgres_db" # use "localhost" for local development, `postgres_db` for docker
  port: "5432"
  log_mode: true # SQL logger , false in production

```

### Develop Application in Docker Compose with Live Reload
Follow these steps:
- Make sure install the latest version of docker and docker-compose
- Docker Installation for your desire OS https://docs.docker.com/engine/install/ubuntu/
- Docker Composer Installation https://docs.docker.com/compose/install/
- Run `make dev`

### Local Setup Instruction
Follow these steps:
- Configuration manage from [config.yml](config.yml) file
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally run `go run main.go` or `go build main.go` and run `./main`

### Routes
- The application available and check health on [0.0.0.0:8000/health](http://0.0.0.0:8000/health)

### Middlewares
- Use Gin CORSMiddleware
```go
router := gin.New()
router.Use(gin.Logger())
router.Use(gin.Recovery())
router.Use(middleware.CORSMiddleware())
```
### Logging
- Use [logrus](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
- `INFO 2022-03-12T00:33:32+03:00 Server is starting at 127.0.0.1:8000`

### Boilerplate Structure
<pre>├── config.yml
├── <font color="#3465A4"><b>controllers</b></font>
│   ├── controller.go
│   └── example_controller.go
├── docker-compose-dev.yml
├── docker-compose-prod.yml
├── Dockerfile
├── Dockerfile-dev
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── Makefile
├── <font color="#3465A4"><b>models</b></font>
│   └── example_model.go
├── <font color="#3465A4"><b>pkg</b></font>
│   ├── <font color="#3465A4"><b>config</b></font>
│   │   ├── config.go
│   │   ├── db.go
│   │   └── server.go
│   ├── <font color="#3465A4"><b>database</b></font>
│   │   ├── database.go
│   │   └── migration.go
│   ├── <font color="#3465A4"><b>helpers</b></font>
│   │   ├── <font color="#3465A4"><b>pagination</b></font>
│   │   │   └── pagination.go
│   │   ├── response.go
│   │   └── search.go
│   ├── <font color="#3465A4"><b>logger</b></font>
│   │   └── logger.go
│   └── <font color="#3465A4"><b>routers</b></font>
│       ├── example.go
│       ├── index.go
│       ├── <font color="#3465A4"><b>middleware</b></font>
│       │   └── cors.go
│       └── router.go
├── README.md
</pre>

### Lets Build an API

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
package database

import (
	"gin-boilerplate/models"
)

//Add list of model add for migrations
var migrationModels = []interface{}{&models.Example{}}

```
3. [controller](controllers) folder add a file `example_controller.go`
- Create API Endpoint 
- Use any syntax of GORM after `base.DB`, this is wrapper of `*gorm.DB`

```go
package controllers

import (
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (base *Controller) CreateExample(ctx *gin.Context) {
	example := new(models.Example)

	err := ctx.ShouldBindJSON(&example)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = base.DB.Create(&example).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &example)
}
```
4. [routers](pkg/routers) folder add a file `example.go`
```go
package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gin-boilerplate/controllers"
)


func TestRoutes(route *gin.Engine, db *gorm.DB) {
	ctrl := controllers.Controller{DB: db}
	v1 := route.Group("/v1")
	v1.POST("/example/", ctrl.CreateExample)
}
```
5. Finally, register routes to [index.go](pkg/routers/index.go)
```go
package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

//RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(route *gin.Engine, db *gorm.DB) {
	//Add All route
	TestRoutes(route, db)
}
```
- Congratulation, your endpoint is ready `0.0.0.0:8000/v1/example/`

### Code Examples
- [Example](examples) contains sample code of different type of example

### Useful Commands

- `make dev`: make dev for development work
- `make build`: make build container
- `make production`: docker production build and up
- `clean`: clean for all clear docker images

### Container Development Build
- Run `make build`

### Container Production Build and Up
- Run `make production`

### Use Packages
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps (Docker Development)
