package main

import (
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("")))
	r.PathPrefix("/views/").Handler(http.FileServer(http.Dir("")))//todo THIS MAY NOT BE SECURE

	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))
}
