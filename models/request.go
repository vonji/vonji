package models

import "github.com/jinzhu/gorm"

type Request struct {
	gorm.Model
	Post
	Title string
	Responses []Response
	Views uint
}
