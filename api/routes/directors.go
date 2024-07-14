package routes

import (
	"github.com/gorilla/mux"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang-BD/api/controllers"
)

func Directors(router *mux.Router) {
	router.HandleFunc("/diretores", controllers.GetDirectors).Methods("GET")
	router.HandleFunc("/diretor/{id}", controllers.GetDirectorById).Methods("GET")
	router.HandleFunc("/diretores", controllers.CreateDirector).Methods("POST")
}
