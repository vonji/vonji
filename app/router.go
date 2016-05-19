package vonji

import (
	"net/http"
	"github.com/gorilla/mux"
)

//TODO maybe use https://github.com/husobee/vestigo for routes
//TODO auto register default routes (send type in param or something)
func RegisterRoutes(r *mux.Router) {
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("")))
	r.PathPrefix("/views/").Handler(http.FileServer(http.Dir("")))//todo THIS MAY NOT BE SECURE
}