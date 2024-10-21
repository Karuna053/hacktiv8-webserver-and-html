package routes

import (
	"webserver-and-html/config"
	"webserver-and-html/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Route for login page.
	router.GET("/login", controller.LoginIndex)
	router.POST("/login", controller.AuthenticateUser)

	return router
}

func RunServer() {
	router := Router()
	router.Run(config.Port)
}
