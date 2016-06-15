package models

import "github.com/jinzhu/gorm"

type Response struct {
	gorm.Model
	Post
	RequestID uint
}
