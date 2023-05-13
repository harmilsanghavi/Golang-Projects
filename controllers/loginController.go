package controllers

import (
	"firstproject/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// calling login form
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// click on login button than it call
func SaveLogin(c *gin.Context) {
	fmt.Println("login")

	Show := services.GetAllData()
	c.HTML(http.StatusOK, "show.html", gin.H{
		"Books": Show,
	})
}
