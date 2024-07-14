package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang-BD/api/models"
	db "github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang-BD/api/services"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang-BD/api/utils"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	var reviews []models.Review
	rows := db.Query("SELECT r.id, m.title AS movie, r.reviewer, r.rating, r.comment FROM review r INNER JOIN movie m ON r.movie = m.id ORDER BY r.id")

	for rows.Next() {
		var review models.Review
		err := rows.StructScan(&review)
		if err != nil {
			log.Fatalln(err)
		}
		reviews = append(reviews, review)
	}

	utils.ConsoleLog(w, r)
	log.Printf("%d Registros retornados", len(reviews))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func GetReviewById(w http.ResponseWriter, r *http.Request) {
	utils.ConsoleLog(w, r)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var review models.Review
	rows := db.Query(fmt.Sprintf("SELECT r.id, m.title AS movie, r.reviewer, r.rating, r.comment FROM review r INNER JOIN movie m ON r.movie = m.id WHERE r.id=%s ORDER BY r.id", params["id"]))
	for rows.Next() {
		err := rows.StructScan(&review)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if review.ID == "" {
		NotFound(w, r, fmt.Sprintf("Nenhum registro para ID=%s encontrado", params["id"]))
		log.Printf("Nenhum registro para ID=%s encontrado", params["id"])
		return
	}

	json.NewEncoder(w).Encode(&review)
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)

	var review models.Review

	err := json.Unmarshal(data, &review)
	utils.ConsoleLog(w, r)

	if err != nil {
		log.Fatalf("Erro ao converter JSON para struct: %s", err)
	}

	rows := db.Query(fmt.Sprintf("INSERT INTO review (movie,reviewer,rating,comment) VALUES (%s,'%s',%d,'%s') RETURNING id", review.Movie, review.Reviewer, review.Rating, review.Comment))

	for rows.Next() {

		err := rows.Scan(&review.ID)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("Análise Inserida: ID=%s MovieID=%s Reviewer=%s Rating=%d Comment=%s", review.ID, review.Movie, review.Reviewer, review.Rating, review.Comment)

	json.NewEncoder(w).Encode(&review)
}
