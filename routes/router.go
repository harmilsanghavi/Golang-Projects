package routes

import (
	"firstproject/controllers"
	"firstproject/middleware"

	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
	router.GET("/", controllers.GetRegister)
	router.GET("/update/:ajax", controllers.GetRegister)
	router.POST("/saveregister", controllers.SaveRegister)
	router.GET("/login", controllers.GetLogin)
	router.GET("/page", controllers.PageSet)
	// router.GET("/page2", controllers.PageDemo)
	router.GET("/delete/:id", controllers.DeleteData)
	router.GET("/emailCheck/:var", controllers.GetEmail)
	router.GET("/logout", controllers.Clear)
	router.POST("/savelogin", middleware.Check(), controllers.PageDemo)
	router.GET("/display", func(c *gin.Context) {
		c.HTML(200, "showData.html", nil)
	})
	router.GET("/emailotpget", controllers.EmailOtpGet)
	router.POST("/getotp", controllers.GetOtp)
	router.GET("/forgetpass", controllers.ForgetPasswordGet)
	router.POST("/setpass", controllers.SetPassword)
	router.GET("/sendtwillio", controllers.GetFileSMS)
	router.POST("/receivesms", controllers.ReceiveSMS)
	router.GET("/sendsms", controllers.SendSMS)
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(302, "/login")
	})
}
