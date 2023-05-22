package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"qianyu/openstage/config"
	"qianyu/openstage/utils"
)

var (
	server *gin.Engine
)

const defaultPort = "8080"

func init() {
	config.ConnectMongodb()
	server = gin.Default()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	defer config.Mongoclient.Disconnect(config.Ctx)

	// middlewares
	// server.Use(middlewares.AuthMiddleware)

	// graphql
	server.POST("/query", utils.GraphqlHandler())
	server.GET("/", utils.PlaygroundHandler())

	basepath := server.Group("/v1")
	config.Uc.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(port))
}
