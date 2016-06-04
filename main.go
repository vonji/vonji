package main

import (
	"net/http"
	"os"

	"github.com/vonji/vonji-api/api"
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/routers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/rs/cors"
)

func main() {
	db, err := gorm.Open("postgres", "user=api password=NOT0 dbname=vonji sslmode=disable")

	db.LogMode(true)

	defer db.Close()
	db.AutoMigrate(&models.Tag{}, &models.User{})

	if err != nil {
		panic(err.Error())
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		Debug:          true,
	})

	r := mux.NewRouter()

	app := api.App{}

	app.Init(r)
	api.InitContext(&app, db)
	routers.RegisterRoutes(r)

	//TODO use something like Alice to chain middlewares
	http.ListenAndServe(":1618", handlers.LoggingHandler(os.Stdout, c.Handler(r)))
}
