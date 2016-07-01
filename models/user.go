package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email		string
	Password	string
	FirstName   string
	LastName    string
	Description string
	Location    string
	VCoins      int
	VReputation int
	Tags        []Tag `gorm:"many2many:user_tags;"` //todo probalement un score par tag
}
