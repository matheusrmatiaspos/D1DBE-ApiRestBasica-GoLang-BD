package server

import (
	"github.com/gorilla/mux"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang-BD/api/routes"
)

func Start(router *mux.Router, port int) {
	routes.Root(router)

	routes.Movies(router)
	routes.Directors(router)
	routes.Actors(router)
	routes.Genres(router)
	routes.Reviews(router)
}
