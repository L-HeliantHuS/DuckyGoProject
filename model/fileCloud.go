package model

import "github.com/jinzhu/gorm"

type FileCloud struct {
	gorm.Model
	FileName string
	Size     int
}
