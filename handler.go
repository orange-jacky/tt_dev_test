package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Binding from json
type UserInputParam struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type RelationshipInputParam struct {
	State string `form:"state" json:"state" binding:"required"`
}

func ListAllUsers(c *gin.Context) {
	users, err := DbListAllUsers()
	if err == nil {
		c.JSON(http.StatusOK, users)
	} else {
		fail := fmt.Sprintf("fail. %s", err.Error())
		c.JSON(http.StatusOK, gin.H{"status": fail})
	}
}

func CreateNewUser(c *gin.Context) {
	var user UserInputParam
	if c.BindJSON(&user) == nil {
		user, err := DbCreateNewUser(user.Name)
		if err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			fail := fmt.Sprintf("fail. %s", err.Error())
			c.JSON(http.StatusOK, gin.H{"status": fail})
		}
	}
}

func ListUserAllRelationships(c *gin.Context) {
	user_id := c.Param("user_id")
	r, err := DbListUserAllRelationships(StrToInt64(user_id))
	if err == nil {
		c.JSON(http.StatusOK, r)
	} else {
		fail := fmt.Sprintf("fail. %s", err.Error())
		c.JSON(http.StatusOK, gin.H{"status": fail})
	}
}

func CreatetUserRelationshipState(c *gin.Context) {
	user_id := c.Param("user_id")
	other_user_id := c.Param("other_user_id")
	var relationship RelationshipInputParam
	if c.BindJSON(&relationship) == nil {
		if relationship.State != "liked" && relationship.State != "disliked" {
			c.String(http.StatusOK, "state=liked|disliked")
			return
		}
		r, err := DbInsertUserRelationshipState(StrToInt64(user_id),
			StrToInt64(other_user_id), relationship.State)
		if err == nil {
			c.JSON(http.StatusOK, r)
		} else {
			fail := fmt.Sprintf("fail. %s", err.Error())
			c.JSON(http.StatusOK, gin.H{"status": fail})
		}
	}
}
