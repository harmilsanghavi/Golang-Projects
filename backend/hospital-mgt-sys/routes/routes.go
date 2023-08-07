package routes

import (
	"hospital-mgt-sys/auth"
	"hospital-mgt-sys/controllers"
	"hospital-mgt-sys/middleware"

	"github.com/gin-gonic/gin"
)

func ManageApi(routes *gin.Engine) {
	routes.POST("/signup", controllers.Signup)
	routes.POST("/signin", controllers.Signin)
	routes.GET("/getdepartment", controllers.GetAlldepartment)
	// routes.GET("/page/:current/:size", controllers.PerPage)
	routes.GET("/refreshtoken", auth.Refresh)
	v1 := routes.Group("/patient")
	v1.Use(middleware.Check())
	{
		v1.GET("/data", controllers.Patientdashboard)
		v1.GET("/doctordetail", controllers.PreDetailAppoinment)
		v1.GET("/checkDate/:dateuser/:id", controllers.Check)
		// v1.GET("/getdoctor", controllers.PreDetailAppoinment)
		v1.POST("/bookappoinment", controllers.AddAppointment)
	}
	v2 := routes.Group("/doctor")
	v2.Use(middleware.Check())
	{
		v2.GET("/data", controllers.Patientdashboard)
		v2.GET("/appoinmentrequest", controllers.AppoinmentRequest)
	}
}
