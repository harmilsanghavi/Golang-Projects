package main

import (
	"fmt"
	"hospital-mgt-sys/controllers"
	"hospital-mgt-sys/initializers"
	"hospital-mgt-sys/routes"
	"hospital-mgt-sys/services"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Loadenv()
	initializers.Conn()
	services.TableCreate()
	controllers.StartWorkerPool(5)
	services.FirstInsertData()
}

func main() {
	fmt.Println("hospital-mgt-sys")
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, bearer")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	routes.ManageApi(router)
	router.Run(":" + os.Getenv("PORT"))
}
