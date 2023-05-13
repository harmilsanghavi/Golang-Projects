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
	initilizers.DB.Find(&Store)
	return Store
}
