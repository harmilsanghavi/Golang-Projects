package controllers

import (
	"hospital-mgt-sys/services"

	"github.com/gin-gonic/gin"
)

type Departments struct {
	Department_name string
	Fname           string
	Lname           string
	Id              int
}

func GetAlldepartment(c *gin.Context) {
	var dep []Departments
	departmentData := services.GetAllDepartment()
	// fmt.Println("all Data :- ", departmentData)
	for _, dep_name := range departmentData {
		dep = append(dep, Departments{
			Department_name: dep_name.Department,
			Fname:           dep_name.First_name,
			Lname:           dep_name.Last_name,
			Id:              dep_name.User_id,
		})
	}
	c.JSON(200, dep)
}
