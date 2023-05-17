package main

import (
	"qianyu/openstage/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.GetRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}