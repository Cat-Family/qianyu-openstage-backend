package main

import (
	"fmt"
	"log"
	"os"
	"qianyu/openstage/config"
	"qianyu/openstage/middlewares"
	"qianyu/openstage/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	config.ConnectMongodb()
	server = gin.Default()

	prvKey, err := os.ReadFile("cert/qianyu-openstage-jwt_rsa")
	if err != nil {
		log.Fatalln(err)
	}
	pubKey, err := os.ReadFile("cert/qianyu-openstage-jwt_rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}

	jwtToken := utils.NewJWT(prvKey, pubKey)
	tok, err := jwtToken.Create(time.Hour, "Can be anything")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("TOKEN:", tok)
}

func main() {
	defer config.Mongoclient.Disconnect(config.Ctx)

	server.Use(middlewares.AuthMiddleware)
	basepath := server.Group("/v1")
	config.Uc.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}
