package main

import (
	"github.com/gin-gonic/gin"
)

func RunGin() {
	router := gin.Default()

	router.GET("/users", ListAllUsers)
	router.POST("/users", CreateNewUser)
	router.GET("/users/:user_id/relationships", ListUserAllRelationships)
	router.PUT("/users/:user_id/relationships/:other_user_id", CreatetUserRelationshipState)

	router.Run(":8080")
}

func main() {
	RunGin()
}
