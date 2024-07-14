package routes

import (
	"github.com/gorilla/mux"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang-BD/api/controllers"
)

func Reviews(router *mux.Router) {
	router.HandleFunc("/analises", controllers.GetReviews).Methods("GET")
	router.HandleFunc("/analise/{id}", controllers.GetReviewById).Methods("GET")
	router.HandleFunc("/analises", controllers.CreateReview).Methods("POST")
}
