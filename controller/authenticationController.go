package controller

import (
	"net/http"
	"webserver-and-html/entity"

	"github.com/gin-gonic/gin"
)

func LoginIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func AuthenticateUser(c *gin.Context) {
	var matchedUser *entity.User
	username := c.PostForm("username")
	password := c.PostForm("password")
	userFound := false

	// Validate user.
	var form entity.User
	if err := c.ShouldBind(&form); err != nil {
		// If validation fails, return an error message
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check if user exists.
	for _, user := range entity.Users {
		if username == user.Username {
			matchedUser = &user
			userFound = true
			break
		}
	}

	if !userFound {
		c.HTML(http.StatusOK, "error-page.html", gin.H{
			"error_message": "User not found.",
		})
		return
	}

	// If password matches.
	if entity.CheckPasswordHash(password, matchedUser.HashedPassword) {
		c.HTML(http.StatusOK, "userdata.html", gin.H{
			"username": matchedUser.Username,
			"email":    matchedUser.Email, // Pass additional fields as needed
		})
	} else {
		c.HTML(http.StatusOK, "error-page.html", gin.H{
			"error_message": "Wrong password.",
		})
	}
}
