package routes

import (
	"github.com/gorilla/mux"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang-BD/api/controllers"
)

func Actors(router *mux.Router) {
	router.HandleFunc("/atores", controllers.GetActors).Methods("GET")
	router.HandleFunc("/ator/{id}", controllers.GetActorsByID).Methods("GET")
	router.HandleFunc("/atores", controllers.CreateActor).Methods("POST")
}
