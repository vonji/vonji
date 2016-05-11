package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"vonji/controllers"
	"github.com/gorilla/handlers"
	"os"
	"vonji/App"
)

func main() {

	db, err := gorm.Open("postgres", "user=api password=NOT0 dbname=vonji sslmode=disable")

	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	r := mux.NewRouter()

	app := vonji.App{}

	app.RegisterController(&controllers.UserController{})

	app.Init(r)
	vonji.Init(&app, db)


	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))
}
