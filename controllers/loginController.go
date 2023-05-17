package controllers

import (
	"firstproject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaginationData struct {
	Next       int
	Current    int
	Previous   int
	TotalPages int
}

// calling login form
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// click on login button than it call
// func SaveLogin(c *gin.Context) {
// 	//fmt.Println("login")
// 	//Get Page Number
// 	perPage := 3
// 	page := 1
// 	pageStr := c.Param("page")

// 	if pageStr != "" {
// 		page, _ = strconv.Atoi(pageStr)
// 	}
// 	//Calculating Offset
// 	Offset := (page - 1) * perPage

// 	Show, totalRows := services.GetAllData(Offset, perPage)
// 	//fmt.Println("total Rows :---", totalRows)

// 	//fmt.Println("demo:::-====", totalRows/int64(perPage))

// 	totalPages := math.Ceil(float64(totalRows / int64(perPage)))

// 	//fmt.Println("total Pages :---", totalPages)

// 	c.HTML(http.StatusOK, "show.html", gin.H{
// 		"Books": Show,
// 		"Pagination": PaginationData{
// 			Next:       page + 1,
// 			Current:    page,
// 			Previous:   page - 1,
// 			TotalPages: int(totalPages),
// 		},
// 	})
// }

func PageSet(c *gin.Context) {

	Store := services.GetAllData()

	c.JSON(200, gin.H{
		"data": Store,
	})
}
func PageDemo(c *gin.Context) {
	c.HTML(200, "showData.html", nil)
}
func Clear(c *gin.Context) {
	c.SetCookie("mycookie", "", -1, "/", "localhost", false, true)
	c.Redirect(302, "/login")
}
