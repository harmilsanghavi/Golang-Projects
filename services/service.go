package services

import (
	"firstproject/initilizers"
	"firstproject/models"
)

var Store []models.Demo

// var Show []models.Show

func Get2data() []models.Demo {
	initilizers.DB.Select("email", "password").Find(&Store)
	return Store
}
func GetAllData() []models.Demo {
	initilizers.DB.Where("is_delete IS NULL OR is_delete=0").Find(&Store)
	return Store
}

func DeleteDataID(id int) {
	initilizers.DB.Model(&models.Demo{}).Where("id = ?", id).Update("is_delete", 1)
}
func AllDataOneUser(id int) []models.Demo {
	initilizers.DB.Where("id=?", id).Find(&Store)
	return Store
}

func UpdateData(id, number int, fname, lname, email, profile string) {
	initilizers.DB.Model(&models.Demo{}).Where("id=?", id).Updates(map[string]interface{}{
		"Fname":   fname,
		"Lname":   lname,
		"Email":   email,
		"Number":  number,
		"Profile": profile,
	})
}
func GetEmail() []models.Demo {
	initilizers.DB.Select("email").Find(&Store)
	return Store
}
