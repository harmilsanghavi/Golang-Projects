package services

import (
	"to-do-list/initilizer"
	"to-do-list/model"
)

var Data []model.Todo

func GetAllData() []model.Todo {
	// initilizer.DB.Find(&Data)
	// initilizer.DB.Model(&model.Todo{}).Preload("Attachments").Find(&Data)
	// return Data

	// if err := initilizer.DB.Preload("Attechments").Find(&Data).Error; err != nil {
	// 	// Handle error
	// 	fmt.Println("error in fetch :-", err)
	// }
	// return Data

	// Retrieve todos with their attachments
	if err := initilizer.DB.Preload("Attachments").Find(&Data).Error; err != nil {

		return nil
	}

	return Data

}

func UpdateData(id int, value string) {
	initilizer.DB.Model(&model.Todo{}).Where("id=?", id).Update("stage", value)
}

func GetUpdatedData(id int) []model.Todo {
	initilizer.DB.Where("id=?", id).Preload("Attachments").Find(&Data)
	return Data
}

// func DeleteTask(id int) {
// 	var user model.Todo
// 	initilizer.DB.Preload("Attachments").First(&user, id)
// 	initilizer.DB.Delete(&user)
// }
