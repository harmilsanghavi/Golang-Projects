package initilizers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// creating gloable variable of *gorm.DB
var DB *gorm.DB

// connecting with database
func DbCoonect() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("cann't conenct to database")
	}
}
