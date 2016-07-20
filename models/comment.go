package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Post
	RequestID uint//TODO better
	ResponseID uint//
}