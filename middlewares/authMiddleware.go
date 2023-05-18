package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"qianyu/openstage/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Define a custom claims struct
type UserClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

// @title AuthMiddleware
// @description a middleware to check if the request is authorized
// @param c *gin.Context
func AuthMiddleware(c *gin.Context) {
	prvKey, err := os.ReadFile("cert/qianyu-openstage-jwt_rsa")
	if err != nil {
		log.Fatalln(err)
	}
	pubKey, err := os.ReadFile("cert/qianyu-openstage-jwt_rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}

	jwtToken := utils.NewJWT(prvKey, pubKey)

	// load jwt secret from .env file
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal(envErr)
	}

	// Get the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Split the header value into its parts
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header format"})
		return
	}

	// Parse the token
	tokenString := parts[1]
	content, err := jwtToken.Validate(tokenString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("CONTENT:", content)

	// Set the claims in the Gin context for later use
	c.Set("claims", content)
}
