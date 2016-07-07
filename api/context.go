package api

import "github.com/jinzhu/gorm"

type Context struct {
	App *App
	DB  *gorm.DB
}

var context = Context{}

func InitContext(app *App, db *gorm.DB) {
	context.App = app
	context.DB = db
}

func GetContext() *Context {
	return &context
}
