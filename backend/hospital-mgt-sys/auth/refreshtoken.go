package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Refresh(c *gin.Context) {
	token := c.GetHeader("Bearer")
	fmt.Println("from middleware", token)
	claims, err := ValidateToken(token)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	loginResponse := SetJwt(claims.Id, claims.Email, claims.Password, claims.Role_name)
	c.JSON(200, gin.H{
		"status":       loginResponse.Success,
		"token":        loginResponse.Token,
		"refreshToken": "",
		"data":         "",
		"message":      "doctor data",
		"toast":        false,
		"id":           claims.Id,
		"role":         claims.Role_name,
	})
}
