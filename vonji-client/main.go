package main

import (
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
	"io"
	"strings"
)

//TODO better name
func load(file string, w io.Writer) error {
	view, err := ioutil.ReadFile(file)

	if err == nil {
		w.Write(view)
	}
	return err
}

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	var path string

	i := strings.Index(r.RequestURI, "?")
	if i > 0 {
		path = r.RequestURI[0:i]
	} else {
		path = r.RequestURI
	}
	if path[0] == '/' {
		path = path[1:len(path)]
	}

	load("views/_header.html", w)
	load(path, w)
	load("views/_footer.html", w)
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("")))
	//	r.PathPrefix("/views/").Handler(http.FileServer(http.Dir("")))//todo THIS MAY NOT BE SECURE
	r.PathPrefix("/views/").HandlerFunc(serveHTTP)

	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))
}
