package vonji

import (
	"github.com/gorilla/mux"
)

type App struct {
	Router      *mux.Router
}

func (a *App) Init(router *mux.Router) {
	a.Router = router
	RegisterRoutes(router)
}