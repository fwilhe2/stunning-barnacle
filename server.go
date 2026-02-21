package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port int
	Secret string
}

func SetupRouter(port int) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/token", func(c *gin.Context) {
		token := createToken("test-user", "my-secret")
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	return r
}

func createToken(user string, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}