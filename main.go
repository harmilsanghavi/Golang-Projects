package main

import (
	"firstproject/initilizers"
	"firstproject/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initilizers.LoadEnvFile()
	initilizers.DbCoonect()
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/public", "./public")
	router.Static("/asset", "./asset")
	routes.Api(router)
	router.Run(":" + os.Getenv("PORT"))
}
