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
func Check() gin.HandlerFunc {
	return func(c *gin.Context) {

		//getting user input from form using post method
		email := c.PostForm("email")
		password := c.PostForm("password")

		fmt.Println(email)
		fmt.Println(password)

		if len(email) <= 0 || len(password) <= 0 {
			c.Abort()
			c.Redirect(302, "/login")
			return
		}

		Store := services.Get2data()

		for _, name := range Store {
			err := bcrypt.CompareHashAndPassword([]byte(name.Password), []byte(password))

			if email == name.Email && err == nil {
				f = 1
				loginResponse = auth.SetJwt(name.Email, name.Password)
				cookie, err := c.Cookie("mycookie")
				if err != nil {
					cookie = "not Set"
					c.SetCookie("mycookie", loginResponse.Token, 3600, "/", "localhost", false, true)
				}
				fmt.Println(cookie)
				break
			} else {

				f = 0
				continue
			}
		}

		if f == 1 {
			c.Set("LoginResponse", loginResponse)
			cookie, err := c.Cookie("mycookie")
			fmt.Println(cookie)
			if err != nil {
				fmt.Println("err in cookie")
				// c.Abort()
				c.Redirect(302, "/login")
			}
			c.Next()
		} else {
			c.Abort()
			c.Redirect(302, "/login")
			return
		}
	}
}
