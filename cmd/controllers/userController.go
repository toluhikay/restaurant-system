package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	fmt.Printf("Jah is Great")
	return func(c *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) string {
	return password
}

func VerifyPassword(userPassword string, inputedPassword string) bool {
	err := false
	if err {
		return err
	}

	return true
}
