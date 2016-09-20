package models

import "github.com/jinzhu/gorm"

type Request struct {
	gorm.Model
	Post
	Title     string
	Responses []Response
	Views     uint
	Tags      []Tag `gorm:"many2many:request_tags;"`
	Status    string
	Duration  uint
	Frequency uint
	FrequencyUnit string
	PeriodStart   string
	PeriodEnd     string
	Location      string
}
