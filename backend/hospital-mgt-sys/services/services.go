package services

import (
	"fmt"
	"hospital-mgt-sys/initializers"
	"hospital-mgt-sys/model"
	"time"
)

var Signups []model.Signup
var availableTime []model.AppoinmentTime
var Doctor []model.Doctor_detail
var Appoinmentbookdata []model.AppoinmentBook
var Times []model.AppoinmentTime

type TimeID struct {
	TimeStr string
	ID      int
}
type DoctorData struct {
	Id         int
	First_name string
	Department string
	Day        string
	Time       time.Time
	Status     string
}

func InsertSignup(signup model.Signup) (uint, error) {
	if err := initializers.DB.Create(&signup).Error; err != nil {
		return 0, err
	}
	return signup.ID, nil
}

func InsertDoctorDetails(doctor model.Doctor_detail) {
	initializers.DB.Create(&doctor)
}

func TableCreate() {
	initializers.DB.AutoMigrate(&model.Doctor_detail{})
	initializers.DB.AutoMigrate(&model.Signup{})
	initializers.DB.AutoMigrate(&model.AppoinmentBook{})
}

func GetUser() []model.Signup {
	initializers.DB.Find(&Signups)
	return Signups
}

func InsertAppoinmentData(data *model.AppoinmentBook) {
	initializers.DB.Create(&data)
}

func GetAllDoctor() []model.Doctor_detail {
	if err := initializers.DB.Find(&Doctor).Error; err != nil {
		return nil
	}
	return Doctor
}

func GetAllDepartment() []model.Doctor_detail {
	var dep []model.Doctor_detail
	initializers.DB.Select("DISTINCT department, first_name, last_name,id").Find(&dep)
	return dep
}

func GetAllDoctor2(id int) []model.Doctor_detail {
	if err := initializers.DB.Where("User_id=?", id).Find(&Doctor).Error; err != nil {
		return nil
	}
	return Doctor
}

func GetBookedTime(doctor_id int, userSelectedDate string) []model.AppoinmentTime {
	var availableTime []model.AppoinmentTime
	initializers.DB.Joins("INNER JOIN appoinment_books ON appoinment_times.id = appoinment_books.apooinment_time").
		Where("appoinment_books.doctor_id = ? AND appoinment_books.day=?", doctor_id, userSelectedDate).
		Select("appoinment_times.time").
		Find(&availableTime)
	return availableTime
}

func GetAllTime(formattedTime, userSelectedDate string) []model.AppoinmentTime {
	currentTime := time.Now()
	formatted := currentTime.Format("2006-01-02")
	layout := "2006-01-02"
	date1Str := formatted
	date2Str := userSelectedDate

	date1, err := time.Parse(layout, date1Str)
	if err != nil {
		panic(err)
	}

	date2, err := time.Parse(layout, date2Str)
	if err != nil {
		panic(err)
	}

	if date1.Before(date2) {
		initializers.DB.Find(&availableTime)
	} else {
		if err := initializers.DB.Where("TIME(time) > ?", formattedTime).Find(&availableTime).Error; err != nil {
			return nil
		}
	}
	return availableTime
}

func GetAppoinmentBookData(id, offset, limit int, Role, column, dir, search string) ([]DoctorData, int64, int64, error) {

	var result []DoctorData

	var count int64
	var totalCount int64
	fmt.Println("search :-", search)
	if Role == "patient" {
		initializers.DB.Table("doctor_details").
			Select("appoinment_books.id,doctor_details.first_name,doctor_details.department,appoinment_books.day,appoinment_times.time,appoinment_books.status").
			Joins("inner join appoinment_books on doctor_details.user_id = appoinment_books.doctor_id").
			Joins("inner join appoinment_times on appoinment_times.id = appoinment_books.apooinment_time").
			Where("appoinment_books.user_id = ? AND appoinment_books.deleted_at IS NULL AND  (doctor_details.first_name LIKE ? OR doctor_details.department LIKE ? OR appoinment_books.day LIKE ? OR appoinment_times.time LIKE ?)", id, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
			Order(column + " " + dir).Offset(offset).
			Limit(limit).
			Find(&result).Count(&count)
		initializers.DB.Table("doctor_details").
			Select("doctor_details.first_name,doctor_details.department,appoinment_books.day,appoinment_times.time,appoinment_books.status").
			Joins("inner join appoinment_books on doctor_details.user_id = appoinment_books.doctor_id").
			Joins("inner join appoinment_times on appoinment_times.id = appoinment_books.apooinment_time").
			Where("appoinment_books.user_id = ? AND appoinment_books.deleted_at IS NULL AND  (doctor_details.first_name LIKE ? OR doctor_details.department LIKE ? OR appoinment_books.day LIKE ? OR appoinment_times.time LIKE ?)", id, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
			Count(&totalCount)

	} else {
		initializers.DB.Table("appoinment_books").Select("appoinment_books.first_name, appoinment_books.day, appoinment_times.time,appoinment_books.status,appoinment_books.id").
			Joins("inner join appoinment_times on appoinment_books.apooinment_time=appoinment_times.id").
			Where("appoinment_books.doctor_id = ? AND appoinment_books.deleted_at IS NULL AND (appoinment_books.first_name LIKE ? OR appoinment_books.day LIKE ? or appoinment_times.time LIKE ? or appoinment_books.status LIKE ?)", id, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
			Order(column + " " + dir).Offset(offset).
			Limit(limit).
			Scan(&result).Count(&count)
		initializers.DB.Table("appoinment_books").Select("appoinment_books.first_name, appoinment_books.day, appoinment_times.time,appoinment_books.status").
			Joins("inner join appoinment_times on appoinment_books.apooinment_time=appoinment_times.id").
			Where("appoinment_books.doctor_id = ? AND appoinment_books.deleted_at IS NULL AND  (doctor_details.first_name LIKE ? OR doctor_details.department LIKE ? OR appoinment_books.day LIKE ? OR appoinment_times.time LIKE ?)", id, "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").
			Count(&totalCount)
	}

	return result, count, totalCount, nil
}

func RequestAccept(id int, status string) model.AppoinmentBook {
	var data model.AppoinmentBook
	if status == "accept" {
		initializers.DB.Model(&model.AppoinmentBook{}).Where("id = ?", id).Update("status", "accept")
	} else if status == "reject" {
		initializers.DB.Table("appoinment_books").Where("id=?", id).Update("status", "reject")
	} else {
		initializers.DB.Table("appoinment_books").Where("id=?", id).Update("status", "complete")
	}
	initializers.DB.Model(&model.AppoinmentBook{}).Where("id = ?", id).Find(&data)
	return data
}
