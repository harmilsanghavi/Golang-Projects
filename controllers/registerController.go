package controllers

import (
	"firstproject/initilizers"
	"firstproject/models"
	"firstproject/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// this is get method which load register form it is first loaded
func GetRegister(c *gin.Context) {
	//this method is load html which take 3 argument 1 http status 2 html file 3 data
	_, err := c.Cookie("mycookie")
	if err != nil {
		ajax := c.Param("ajax")
		fmt.Println("error", ajax)
		if ajax == "" {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"data":  nil,
				"page":  0,
				"space": " ",
			})
		} else {
			fmt.Println(ajax)
			id, _ := strconv.Atoi(ajax)
			Store := services.AllDataOneUser(id)
			fmt.Println(Store)
			c.HTML(http.StatusOK, "register.html", gin.H{
				"data": Store,
				"page": id,
			})
		}
	} else {
		ajax := c.Param("ajax")
		if ajax == "" {
			c.HTML(http.StatusOK, "showData.html", nil)
		} else {
			fmt.Println("connect", ajax)
			id, err := strconv.Atoi(ajax)
			fmt.Println("convert Error:---", err)
			if err != nil {
				fmt.Println("inner")
				// c.HTML(302, "/logout", nil)
				c.Redirect(302, "/logout")
			}
			Store := services.AllDataOneUser(id)
			fmt.Println(Store)
			c.HTML(http.StatusOK, "register.html", gin.H{
				"data": Store,
				"page": id,
			})
		}

	}

}

// this method is post method when user click on register button
func SaveRegister(c *gin.Context) {
	//getting all data from post method
	i1d := c.PostForm("id")
	// fmt.Println("id::llllll:::::::", i1d)
	// fmt.Println("trim :---", strings.Trim(i1d, " "))
	id, _ := strconv.Atoi(strings.Trim(i1d, " "))
	fname := c.PostForm("fname")
	lname := c.PostForm("lname")
	email := c.PostForm("email")
	password := c.PostForm("password")

	number1 := c.PostForm("number")
	//converting string to int
	// number, _ := strconv.Atoi(number1)
	number, _ := strconv.ParseInt(number1, 10, 64)

	file, err := c.FormFile("image")
	if err != nil {

		fmt.Println(err)
	}

	// Save the file to a directory on the server

	err = c.SaveUploadedFile(file, "/home/harmil-sanghavi/go/src/firstproject/asset/images/"+file.Filename)
	if err != nil {
		fmt.Println(err)
	}

	if id == 0 {

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
			Number:   int(number),
			Profile:  "/images/" + file.Filename,
		}

		//inserting data into table
		result := initilizers.DB.Create(&Demo2)

		//error checking
		if result.Error != nil {
			fmt.Println(result.Error)
			c.Status(400)
		} else {
			c.Status(200)
			c.Redirect(302, "/login")
		}
	} else {

		profile := "/images/" + file.Filename

		services.UpdateData(id, int(number), fname, lname, email, profile)
		c.Redirect(302, "/display")
	}
}

var f int = 0

func GetEmail(c *gin.Context) {
	value := c.Param("var")
	fmt.Println("::::::::", value)
	Email := services.GetEmail()
	for _, data := range Email {
		if value == data.Email {
			f = 1
			message := "Match"
			c.Set("message", message)
			c.JSON(302, message)
			break
		}
	}
	if f == 0 {
		message := "not Match"
		c.Set("message", message)
		c.JSON(302, message)
	}
}
