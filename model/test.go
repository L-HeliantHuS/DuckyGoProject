package model

import "github.com/jinzhu/gorm"

type Test struct {
	gorm.Model
	Title  string
	Number int
	Alias  string
}


