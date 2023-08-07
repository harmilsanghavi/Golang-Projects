package model

import (
	"time"

	"gorm.io/gorm"
)

type Signup struct {
	gorm.Model
	ID        uint   `gorm:"primary_key"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role_name string `json:"role_name"`
}

type Combine struct {
	User   Signup         `json:"user"`
	Doctor *Doctor_detail `json:"doctor,omitempty"`
}

type AppoinmentTime struct {
	gorm.Model
	Time time.Time
}

type AppoinmentBook struct {
	gorm.Model
	First_name      string `json:"fname"`
	Last_name       string `json:"lname"`
	Address         string `json:"address"`
	Number          int    `json:"number"`
	Gender          string `json:"gender"`
	Age             int    `json:"age"`
	Day             string `json:"date"`
	Apooinment_time string `json:"time"`
	Doctor_id       int    `json:"doctor"`
	Role_name       string
	Status          string `json:"status"`
	User_id         int
	Users           Signup `gorm:"foreignKey:User_id"`
}

type Doctor_detail struct {
	gorm.Model
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Number        int    `json:"number"`
	Gender        string `json:"gender"`
	Age           int    `json:"age"`
	Qualification string `json:"qualification"`
	Department    string `json:"department"`
	User_id       int
	Users         Signup `gorm:"foreignKey:User_id"`
}

type Temp struct {
	gorm.Model
	Time time.Time
}
