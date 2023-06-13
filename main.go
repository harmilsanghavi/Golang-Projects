package main

import (
	"fmt"
	"to-do-list/initilizer"
	"to-do-list/router"
)

func init() {
	initilizer.LoadEnvFile()
	initilizer.ConnectDB()
}

func main() {
	fmt.Println("hello")
	router.Routes()
}
