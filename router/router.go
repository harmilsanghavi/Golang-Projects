package router

import (
	"os"
	"to-do-list/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.LoadHTMLGlob("templet/*")
	router.Static("/public", "./public")
	router.GET("/", controllers.ShowAllList)
	router.POST("/addtodo", controllers.AddToDo)
	router.GET("/getupdateddata", controllers.Hold)
	router.Run(":" + os.Getenv("PORT"))
}
