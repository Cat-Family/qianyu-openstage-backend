package controllers

import "github.com/gin-gonic/gin"

func GetUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "some user info",
	})
}