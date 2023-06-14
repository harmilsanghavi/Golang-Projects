package main

import (
	"fmt"
	"to-do-list/initilizer"
	"to-do-list/model"
	"to-do-list/router"
)

func init() {
	initilizer.LoadEnvFile()
	initilizer.ConnectDB()
	initilizer.DB.AutoMigrate(&model.Attechment{})
}

func main() {
	fmt.Println("hello")
	router.Routes()
}
