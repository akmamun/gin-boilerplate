package main

import (
	"log"
	"net/http"
	env "pkg/src/helpers"
	"pkg/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	host := env.GetEnv("HOST", "0.0.0.0")
	port := env.GetEnv("PORT", "8000")

	// database.DbConnection()

	// gin.SetMode(gin.ReleaseMode) //enable in production

	router := gin.New()
	router.Use(gin.Logger())
	routes.ExampleRoutes(router)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": nil, "message": "system is live "})
	})
	log.Printf("API is running %v:%v\n", host, port)

	router.Run(host + ":" + port)
}
