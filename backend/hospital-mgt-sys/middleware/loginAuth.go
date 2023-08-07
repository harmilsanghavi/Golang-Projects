package middleware

import (
	"fmt"
	"hospital-mgt-sys/auth"

	"github.com/gin-gonic/gin"
)

func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.Request.Header["Bearer"] //it is also right way to get Header
		fmt.Println("middleware")
		token := c.GetHeader("Bearer")
		// fmt.Println("from middleware", token)
		// if strings.HasPrefix(token, "Bearer ") {
		// 	// Extract the token by removing "Bearer " from the header value
		// 	token := token[len("Bearer "):]

		// 	// Now you have the token, and you can use it for further processing
		// 	fmt.Println("Extracted token:", token)
		// } else {
		// 	c.JSON(401, gin.H{
		// 		"error": "Invalid token format",
		// 	})
		// }
		claims, err := auth.ValidateToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("Role", claims.Role_name)
		c.Set("id", claims.Id)
		c.Next()
	}
}
