package controllers

import (
	"fmt"
	"hospital-mgt-sys/model"
	"hospital-mgt-sys/services"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func AddAppointment(c *gin.Context) {
	fmt.Println("hello")
	var Data model.AppoinmentBook

	c.Bind(&Data)
	Data.Status = "pending"

	// fmt.Println("doctor id for booked :- ", Data)
	// fmt.Println("fname", Data.First_name)
	// fmt.Println("lname", Data.Last_name)
	// fmt.Println("number", Data.Number)
	// fmt.Println("address", Data.Address)
	// fmt.Println("age", Data.Age)
	// fmt.Println("time", Data.Apooinment_time)
	// fmt.Println("day", Data.Day)
	// fmt.Println("day", Data.Doctor_id)

	role, _ := c.Get("Role")
	id, _ := c.Get("id")
	fmt.Println("claims", role, id)
	Data.User_id = id.(int)
	Data.Role_name = role.(string)

	services.InsertAppoinmentData(&Data)

	c.JSON(200, gin.H{
		"status":  true,
		"token":   "",
		"data":    Data,
		"message": "Appoinment is booked",
		"toast":   true,
	})
}

type DoctorNeed struct {
	Name       string
	Id         int
	Department string
}

func PreDetailAppoinment(c *gin.Context) {
	var detail []DoctorNeed

	doctor := services.GetAllDoctor()

	for _, doctors := range doctor {
		detail = append(detail, DoctorNeed{
			Name:       doctors.First_name,
			Id:         int(doctors.User_id),
			Department: doctors.Department,
		})
	}

	c.JSON(200, gin.H{
		"status":  true,
		"token":   "",
		"data":    detail,
		"message": "All department Data with doctor data",
		"toast":   false,
	})
}

type TimeID struct {
	TimeStr string
	ID      int
}

type EmailJob struct {
	Message *gomail.Message
}

func Check(c *gin.Context) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05.000")

	userSelectedDate := c.Param("dateuser")
	doctor_id := c.Param("id")

	id, _ := strconv.Atoi(doctor_id)
	data := services.GetBookedTime(id, userSelectedDate)
	allTime := services.GetAllTime(formattedTime, userSelectedDate)
	var available []TimeID
	for _, all := range allTime {
		var f int = 0
		for _, booked := range data {
			if booked.Time.Format("15:04 PM") == all.Time.Format("15:04 PM") {
				f = 1
				continue
			}
		}
		if f == 0 {
			available = append(available, TimeID{
				TimeStr: all.Time.Format("03:04 PM"),
				ID:      int(all.ID),
			})
		}
	}
	// fmt.Println("free time :-", available)
	c.JSON(200, gin.H{
		"status":  true,
		"token":   "",
		"data":    available,
		"message": "Available Time of Particular Doctor",
		"toast":   false,
	})
}

var jobQueue chan EmailJob = make(chan EmailJob)

func worker() {
	for job := range jobQueue {
		// Send email using gomail
		dialer := gomail.NewDialer("smtp.gmail.com", 587, "harmil.sanghavi.23@gmail.com", "ylbimcamjxpoatud")
		if err := dialer.DialAndSend(job.Message); err != nil {
			log.Println("Error sending email:", err)
		}
	}
}
func StartWorkerPool(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go worker()
	}
}
func AppoinmentRequest(c *gin.Context) {

	// Add message to job queue

	// fmt.Println("accept")
	id, _ := strconv.Atoi(c.Query("Id"))
	status := c.Query("Status")
	// fmt.Println("id :=- ", id)
	// fmt.Println("status :=- ", status)

	rsponseData := services.RequestAccept(id, status)

	// fmt.Println("response from query :- ", rsponseData)
	m := gomail.NewMessage()
	m.SetHeader("From", "harmil.sanghavi.23@gmail.com")
	m.SetHeader("To", "test@mailinator.com")
	m.SetHeader("Subject", "Appoinemtn Status")
	m.SetBody("text/plain", `Your Appoinment Booked is now `+status)

	job := EmailJob{Message: m}
	jobQueue <- job
	c.JSON(200, gin.H{
		"status":  true,
		"message": status,
		"data":    rsponseData,
		"toast":   false,
	})
}
