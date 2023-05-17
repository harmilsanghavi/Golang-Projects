package models

import "gorm.io/gorm"

// this is model for database which create table using orm
type Demo struct {
	gorm.Model
	Id       int    `gorm:"primary_key" json:"id"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Number   int    `json:"number"`
	Profile  string `json:"profile"`
}
