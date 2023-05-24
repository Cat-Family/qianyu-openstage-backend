package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"qianyu/openstage/config"
	"qianyu/openstage/middlewares"
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

	// graphql
	server.GET("/", utils.PlaygroundHandler())
	server.POST("/query", middlewares.AuthMiddleware, utils.GraphqlHandler())

	basepath := server.Group("/v1")
	// basepath.Use(middlewares.AuthMiddleware)
	config.Uc.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(port))
}
