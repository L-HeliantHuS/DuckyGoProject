package model

import "github.com/jinzhu/gorm"

// Msg 留言板模型
type Msg struct {
	gorm.Model
	UserID  uint
	Message string
	Subject string
}