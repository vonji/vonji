package models

import "github.com/jinzhu/gorm"

type Achievement struct {
	gorm.Model
	Award uint
	Name string
	Description string
	Category string
	CheckID uint
	CheckData int
}
