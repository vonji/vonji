package models

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	FromID uint
	From User
	ToID uint
	To User
	Reason string
	Source string
	Type string// VCOIN | VACTION
	Amount int
}