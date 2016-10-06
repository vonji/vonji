package models

import "github.com/jinzhu/gorm"

type Notification struct {
	gorm.Model
	UserID  uint
	User    User
	Title   string
	Message string
	Read    bool
}