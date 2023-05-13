package middleware

import (
	"firstproject/auth"
	"firstproject/services"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var f int = 0
var loginResponse *auth.LoginResponse

// this function for login checking that user is allready register or not
func Check(c *gin.Context) {

	//getting user input from form using post method
	email := c.PostForm("email")
	password := c.PostForm("password")

	fmt.Println(email)
	fmt.Println(password)

	//creating the []models.Demo variable to store database value

	//orm query to get all record with only email and password
	Store := services.Get2data()

	//printig all value from the database
	fmt.Println("sore data", Store)

	//extracting email from the value which come from the database

	for _, name := range Store {
		err := bcrypt.CompareHashAndPassword([]byte(name.Password), []byte(password))
		if email == name.Email && err == nil {
			f = 1
			loginResponse = auth.SetJwt(name.Email, name.Password)
			break
		} else {
			fmt.Println(f)
			f = 0
			continue
		}
	}
	if f == 1 {
		fmt.Println(loginResponse)
		c.Set("LoginResponse", loginResponse)
	} else {
		c.Redirect(302, "/")
	}
}
