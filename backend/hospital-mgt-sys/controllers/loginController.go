package controllers

import (
	"fmt"
	"hospital-mgt-sys/auth"
	"hospital-mgt-sys/model"
	"hospital-mgt-sys/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	var Signup model.Signup

	if err := c.ShouldBindJSON(&Signup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("1")
	if len(Signup.Email) <= 0 || len(Signup.Password) <= 0 {
		fmt.Println("2")
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{
			"length": "Not Matched",
		})
		return

	}

	fmt.Println("email", Signup.Email)

	dbData := services.GetUser()
	f := 0
	for _, email := range dbData {
		fmt.Println("user defined :-", Signup.Email)
		fmt.Println("Databased defined :-", email.Email)
		if email.Email == Signup.Email {
			loginResponse := auth.SetJwt(int(email.ID), Signup.Email, Signup.Password, email.Role_name)
			fmt.Println("login Resposne :- ", loginResponse)
			if email.Role_name == "doctor" {
				doctorData := services.GetAllDoctor2(int(email.ID))
				c.JSON(200, gin.H{
					"status":       loginResponse.Success,
					"token":        loginResponse.Token,
					"refreshToken": loginResponse.RefreshToken,
					"data":         doctorData,
					"message":      "doctor data",
					"toast":        false,
					"id":           email.ID,
					"role":         email.Role_name,
				})
			} else {
				c.JSON(200, gin.H{
					"id":           email.ID,
					"status":       loginResponse.Success,
					"token":        loginResponse.Token,
					"refreshToken": loginResponse.RefreshToken,
					"data": gin.H{
						"email": email.Email,
					},
					"message": "patient data",
					"toast":   false,
					"role":    email.Role_name,
				})
			}
			f = 1
			break
		}
	}
	if f == 0 {
		c.JSON(200, gin.H{
			"Success": false,
			"Token":   "",
		})
	}
}
