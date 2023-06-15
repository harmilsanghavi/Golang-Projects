package controllers

import (
	"fmt"
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

// func AddToDo(c *gin.Context) {
// 	var todo_id = c.PostForm("todo_id")

// 	fmt.Println("updated id:=-", todo_id)

// 	var data model.Todo

// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		// handle error
// 		fmt.Println(err)
// 	}

// 	if err := c.Bind(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 		return
// 	}

// 	data.Stage = "todo"

// 	files := form.File["upload"]
// 	attachments := make([]model.Attechment, len(files))

// 	for i, file := range files {

// 		err = c.SaveUploadedFile(file, "/home/harmil-sanghavi/go/src/to-do-list/asset/images/"+file.Filename)
// 		if err != nil {

// 			fmt.Println(err)
// 		}
// 		attachments[i] = model.Attechment{
// 			Files:  "/asset/images/" + file.Filename,
// 			TodoID: data.Id,
// 		}

// 		if err := initilizer.DB.Create(&attachments[i]).Error; err != nil {
// 			fmt.Println("error")
// 		}

// 	}
// 	data.Attachments = attachments
// 	initilizer.DB.AutoMigrate(&model.Todo{})
// 	result := initilizer.DB.Create(&data)

// 	if result.Error != nil {
// 		fmt.Println(result.Error)
// 		c.Status(400)
// 	} else {
// 		c.Status(200)
// 		c.Redirect(302, "/")
// 	}
// }

func AddToDo(c *gin.Context) {
	var todo model.Todo
	todo_id := c.PostForm("todo_id")

	// Check if the todo ID is provided
	if todo_id != "" {
		// Update an existing todo
		initilizer.DB.First(&todo, todo_id)
	}

	// Update the todo fields
	todo.Todotopic = c.PostForm("todotopic")
	todo.Tododescription = c.PostForm("tododescription")
	todo.Stage = "todo"

	// Save the todo record
	initilizer.DB.Save(&todo)

	// Get the attachments from the request
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}

	// Create or update the attachments
	files := form.File["upload"]
	attachments := make([]model.Attechment, len(files))

	for i, file := range files {
		err = c.SaveUploadedFile(file, "/home/harmil-sanghavi/go/src/to-do-list/asset/images/"+file.Filename)
		if err != nil {
			fmt.Println(err)
		}

		// Create or update the attachment record
		attachment := model.Attechment{
			Files:  "/asset/images/" + file.Filename,
			TodoID: todo.Id,
		}
		initilizer.DB.Where(model.Attechment{ID: i + 1}).Assign(attachment).FirstOrCreate(&attachment)
		attachments[i] = attachment
	}

	// Update the attachments for the todo
	todo.Attachments = attachments
	initilizer.DB.Save(&todo)

	// Redirect to the home page
	c.Redirect(302, "/")
}

func Hold(c *gin.Context) {
	fmt.Println("hold")
	id := c.Query("id")
	value := c.Query("value")
	fmt.Println("id :- ", id)
	fmt.Println("value :- ", value)
	num, _ := strconv.ParseInt(id, 10, 64)
	var data []model.Todo

	if value == "" {
		data = services.GetUpdatedData(int(num))
	} else {
		services.UpdateData(int(num), value)
		data = services.GetUpdatedData(int(num))
	}
	c.JSON(302, data)
}

// func RetriveHasheMany(c *gin.Context) {
// 	var todos []model.Todo

// 	// Retrieve todos with their attachments
// 	if err := initilizer.DB.Preload("Attachments").Find(&todos).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todos"})
// 		return
// 	}

// 	// Return the todos as a JSON response
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": todos,
// 	})
// }
