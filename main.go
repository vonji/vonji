package main

import (
	"net/http"
	"os"

	"github.com/vonji/vonji-api/api"
	"github.com/vonji/vonji-api/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/rs/cors"

	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/services"
)

//TODO dependecy injection?
func main() {
	db, err := gorm.Open("postgres", "user=api password=NOT0 dbname=vonji sslmode=disable")

	db.LogMode(true)

	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		Debug:          false,
	})

	r := mux.NewRouter()

	app := api.App{}

	app.Init(r)
	api.InitContext(&app, db)
	routes.RegisterRoutes(r)

	initDB(db)

	//TODO use something like Alice to chain middlewares
	http.ListenAndServe(":1618", handlers.LoggingHandler(os.Stdout, c.Handler(r)))
}

func initDB(db *gorm.DB) {
	db.AutoMigrate(&models.Tag{}, &models.User{}, &models.Request{}, &models.Response{}, &models.Comment{})
	if len(services.User.GetAll()) == 0 {
		services.User.Create(&models.User{
			Email: "admin@vonji.fr",
			Password: "admin",
			FirstName: "Admin",
			LastName: "Admin",
			Description: "THE ALMIGHTY ONE",
		})
	}
}