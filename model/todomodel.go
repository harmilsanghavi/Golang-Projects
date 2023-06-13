package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Id              int
	Todotopic       string `form:"todotopic"`
	Tododescription string `form:"tododescription"`
	Stage           string
}
