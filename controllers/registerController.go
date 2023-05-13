package controllers

import (
	"firstproject/initilizers"
	"firstproject/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// this is get method which load register form it is first loaded
func GetRegister(c *gin.Context) {
	//this method is load html which take 3 argument 1 http status 2 html file 3 data
	c.HTML(http.StatusOK, "register.html", nil)
}

// this method is post method when user click on register button
func SaveRegister(c *gin.Context) {
	//getting all data from post method
	fname := c.PostForm("fname")
	lname := c.PostForm("lname")
	email := c.PostForm("email")
	password := c.PostForm("password")
	// profile := c.PostForm("image")
	number1 := c.PostForm("number")
	//converting string to int
	number, _ := strconv.ParseInt(number1, 10, 64)

	// fmt.Println(fname)
	// fmt.Println(lname)
	// fmt.Println(email)
	// fmt.Println(password)
	// fmt.Println(number)
	// fmt.Println(profile)

	// var bindstructure models.Demo
	// c.Bind(&bindstructure)
	// fmt.Println(bindstructure.Fname)
	file, err := c.FormFile("image")
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println("file Name :----", file)
	// Save the file to a directory on the server

	err = c.SaveUploadedFile(file, "/home/harmil-sanghavi/go/src/firstproject/asset/images/"+file.Filename)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}

	//creting table is not exist
	initilizers.DB.AutoMigrate(&models.Demo{})

	//genreting hash password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	//bind the user data into struct
	Demo2 := models.Demo{
		Fname: fname,
		Lname: lname,
		Email: email,
		//hash password is in byte formate so convert it into string
		Password: string(bytes),
		Number:   number,
		Profile:  "/images/" + file.Filename,
	}

	//inserting data into table
	result := initilizers.DB.Create(&Demo2)

	//error checking
	if result.Error != nil {
		c.Status(400)
	} else {
		c.Status(200)
		c.Redirect(302, "/")
	}
}
