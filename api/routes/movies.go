package routes

import (
	"github.com/gorilla/mux"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang-BD/api/controllers"
)

func Movies(router *mux.Router) {
	router.HandleFunc("/filmes", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/filme/{id}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/filmes", controllers.CreateMovie).Methods("POST")
}
