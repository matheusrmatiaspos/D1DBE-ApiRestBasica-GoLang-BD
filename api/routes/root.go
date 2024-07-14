package routes

import (
	"github.com/gorilla/mux"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang-BD/api/controllers"
)

func Root(router *mux.Router) {
	router.HandleFunc("/", controllers.HandleRoot).Methods("GET")
}
