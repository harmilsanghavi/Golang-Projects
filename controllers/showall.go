package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"to-do-list/initilizer"
	"to-do-list/model"
	"to-do-list/services"

	"github.com/gin-gonic/gin"
)

func ShowAllList(c *gin.Context) {
	data := services.GetAllData()
	if data == nil {
		fmt.Println("inner")
		c.HTML(302, "show.html", gin.H{
			"data": nil,
		})
	} else {
		fmt.Println("outer")
		c.HTML(302, "show.html", gin.H{
			"data": data,
		})
	}
}

func AddToDo(c *gin.Context) {
	var data model.Todo
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	data.Stage = "todo"

	initilizer.DB.AutoMigrate(&model.Todo{})
	result := initilizer.DB.Create(&data)

	if result.Error != nil {
		fmt.Println(result.Error)
		c.Status(400)
	} else {
		c.Status(200)
		c.Redirect(302, "/")
	}
}

func Hold(c *gin.Context) {
	fmt.Println("hold")
	id := c.Query("id")
	value := c.Query("value")
	fmt.Println("id :- ", id)
	fmt.Println("value :- ", value)
	num, _ := strconv.ParseInt(id, 10, 64)
	services.UpdateData(int(num), value)
	data := services.GetUpdatedData(int(num))
	c.JSON(302, data)
}
