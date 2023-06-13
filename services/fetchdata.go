package services

import (
	"to-do-list/initilizer"
	"to-do-list/model"
)

var Data []model.Todo

func GetAllData() []model.Todo {
	initilizer.DB.Find(&Data)
	return Data
}

func UpdateData(id int, value string) {
	initilizer.DB.Model(&model.Todo{}).Where("id=?", id).Update("stage", value)
}

func GetUpdatedData(id int) []model.Todo {
	initilizer.DB.Where("id=?", id).Find(&Data)
	return Data
}
