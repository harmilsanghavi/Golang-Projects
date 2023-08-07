package controllers

import (
	"fmt"
	"hospital-mgt-sys/model"
	"hospital-mgt-sys/services"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var requestData model.Combine
	err := c.BindJSON(&requestData)
	if err != nil {
		// handle error
		panic(err)
	}

	fmt.Println("data in signup :- ", requestData)

	if requestData.User.Role_name == "doctor" {
		fmt.Println("if")
		id, err := services.InsertSignup(requestData.User)
		if err != nil {
			panic(err)
		}
		requestData.Doctor = &model.Doctor_detail{
			First_name:    requestData.Doctor.First_name,
			Last_name:     requestData.Doctor.Last_name,
			Number:        requestData.Doctor.Number,
			Gender:        requestData.Doctor.Gender,
			Age:           requestData.Doctor.Age,
			Qualification: requestData.Doctor.Qualification,
			Department:    requestData.Doctor.Department,
			User_id:       int(id),
		}
		services.InsertDoctorDetails(*requestData.Doctor)
		c.JSON(200, gin.H{
			"data":    "User Singup Data as doctor",
			"status":  true,
			"token":   "",
			"message": "user Singup Successfully",
			"toast":   true,
		})

	} else {
		fmt.Println("else")
		services.InsertSignup(requestData.User)
		requestData.Doctor = nil
		c.JSON(200, gin.H{
			"data":    "User Singup Data as Patient",
			"status":  true,
			"token":   "",
			"message": "user Singup Successfully",
			"toast":   true,
		})
	}
}
