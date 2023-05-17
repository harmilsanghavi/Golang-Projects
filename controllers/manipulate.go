package controllers

import (
	"firstproject/services"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteData(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("delete :---", id)
	id1, _ := strconv.Atoi(id)
	services.DeleteDataID(id1)
	c.Redirect(302, "/page2")
}
