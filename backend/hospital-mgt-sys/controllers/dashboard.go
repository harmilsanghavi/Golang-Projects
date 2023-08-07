package controllers

import (
	"fmt"
	"hospital-mgt-sys/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentRow struct {
	Id         int
	Day        string
	Time       string
	First_name string
	Department string
	Status     string
}

func Patientdashboard(c *gin.Context) {
	id, _ := c.Get("id")
	Role, _ := c.Get("Role")

	page, _ := strconv.Atoi(c.Query("current"))
	size, _ := strconv.Atoi(c.Query("size"))
	column := c.Query("column")

	dir := c.Query("dir")
	search := c.DefaultQuery("search", "%")

	offset := (page - 1) * size
	limit := size

	var data []services.DoctorData
	var count int64
	var totalCount int64
	var err error

	data, count, totalCount, err = services.GetAppoinmentBookData(id.(int), offset, limit, Role.(string), column, dir, search)

	appointmentRows := make([]AppointmentRow, len(data))
	if Role == "patient" {
		for i := 0; i < len(data); i++ {
			appointmentRows[i] = AppointmentRow{
				Id:         data[i].Id,
				Day:        data[i].Day,
				Time:       data[i].Time.Format("03:04 PM"),
				First_name: data[i].First_name,
				Department: data[i].Department,
				Status:     data[i].Status,
			}
		}
	} else {
		for i := 0; i < len(data); i++ {
			appointmentRows[i] = AppointmentRow{
				Id:         data[i].Id,
				Day:        data[i].Day,
				Time:       data[i].Time.Format("03:04 PM"),
				First_name: data[i].First_name,
				Status:     data[i].Status,
			}
		}
	}

	fmt.Println("error value :- ", err)
	if err != nil {
		fmt.Println("error")
		c.JSON(200, gin.H{
			"status":  false,
			"token":   "",
			"data":    "",
			"message": "failed to load data",
			"toast":   true,
		})
	}
	fmt.Println("else")

	c.JSON(200, gin.H{
		"status":     true,
		"count":      int(count),
		"totalCount": int(totalCount),
		"data":       appointmentRows,
		"message":    "Appoinment list",
		"toast":      false,
	})
}
