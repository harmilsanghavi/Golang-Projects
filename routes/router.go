package routes

import (
	"firstproject/controllers"
	"firstproject/middleware"
	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
	router.GET("/", controllers.GetRegister)
	router.POST("/saveregister", controllers.SaveRegister)
	router.GET("/login", controllers.GetLogin)
	router.POST("/savelogin", middleware.Check, controllers.SaveLogin)
}
