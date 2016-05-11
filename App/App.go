package vonji

import (
	"vonji/controllers"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	controllers []controllers.Controller
}

func (a *App) Init(router *mux.Router) {
	a.Router = router
	for _, e := range a.controllers {
		e.Init(a.Router)//todo maybe pass a subrouter instead
	}
}

func (a *App) RegisterController(c controllers.Controller) {
	a.controllers = append(a.controllers, c)
}


