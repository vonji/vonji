package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email         string
	Password      string
	DisplayedName string
	RealName      string
	Description   string
	Motto         string
	FacebookLink  string
	TwitterLink   string
	LinkedInLink  string
	Phone         string
	Birthday      string
	Location      string
	VCoins        int
	VAction       int
	Avatar        string
	Gender	      string
	Achievements  []Achievement `gorm:"many2many:user_achievements;"`
	Tags          []Tag `gorm:"many2many:user_tags;"`
}
