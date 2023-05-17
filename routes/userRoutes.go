package routes

import (
	"qianyu/openstage/controllers"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/user")

	// localhost:8080/v1/user GET
	r.GET("/", controllers.GetUserInfo)
}