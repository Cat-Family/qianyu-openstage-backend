package main

import (
	"log"
	"qianyu/openstage/config"

	"github.com/gin-gonic/gin"
)

var (
	server	*gin.Engine
)


func init() {
	config.ConnectMongodb()
	server = gin.Default()
}

func main() {
	defer config.Mongoclient.Disconnect(config.Ctx)
	
	basepath := server.Group("/v1")
	config.Uc.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}