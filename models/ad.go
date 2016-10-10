package models

import "github.com/jinzhu/gorm"

type Ad struct {
	gorm.Model
	Latitude  float64
	Longitude float64
	Region    string
	Url       string
	ImageUrl  string
	AltText   string
}
