package routes

import (
	"firstproject/controllers"
	"firstproject/middleware"

	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
	router.GET("/:ajax", controllers.GetRegister)
	router.POST("/saveregister", controllers.SaveRegister)
	router.GET("/login", controllers.GetLogin)
	router.GET("/page", controllers.PageSet)
	router.GET("/page2", controllers.PageDemo)
	router.GET("/delete/:id", controllers.DeleteData)
	router.GET("/emailCheck/:var", controllers.GetEmail)
	router.GET("/logout", controllers.Clear)
	router.POST("/savelogin", middleware.Check(), controllers.PageDemo)
}
