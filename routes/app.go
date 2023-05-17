package routes

import "github.com/gin-gonic/gin"


func GetRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	addUserRoutes(v1)

	v2 := r.Group("/v2")
	addUserRoutes(v2)
}