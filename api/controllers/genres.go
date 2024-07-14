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

func GetGenres(w http.ResponseWriter, r *http.Request) {
	var genres []models.Genre
	rows := db.Query("SELECT g.id, g.name, g.description, COUNT(m.genre) AS movie_count, g.created_at FROM genre g INNER JOIN movie m ON g.id=m.genre GROUP BY g.id ORDER BY g.id")

	for rows.Next() {
		var genre models.Genre
		err := rows.StructScan(&genre)
		if err != nil {
			log.Fatalln(err)
		}
		genres = append(genres, genre)
	}

	utils.ConsoleLog(w, r)
	log.Printf("%d Registros retornados", len(genres))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}

func GetGenreById(w http.ResponseWriter, r *http.Request) {
	utils.ConsoleLog(w, r)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var genre models.Genre
	rows := db.Query(fmt.Sprintf("SELECT g.id, g.name, g.description, COUNT(m.genre) AS movie_count, g.created_at FROM genre g LEFT JOIN movie m ON g.id=m.genre WHERE g.id=%s GROUP BY g.id ORDER BY g.id", params["id"]))
	for rows.Next() {
		err := rows.StructScan(&genre)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if genre.ID == "" {
		NotFound(w, r, fmt.Sprintf("Nenhum registro para ID=%s encontrado", params["id"]))
		log.Printf("Nenhum registro para ID=%s encontrado", params["id"])
		return
	}

	json.NewEncoder(w).Encode(&genre)
}

func CreateGenre(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)

	var genre models.Genre

	err := json.Unmarshal(data, &genre)
	utils.ConsoleLog(w, r)

	if err != nil {
		log.Fatalf("Erro ao converter JSON para struct: %s", err)
	}

	rows := db.Query(fmt.Sprintf("INSERT INTO genre (name, description,created_at) VALUES('%s','%s','%s') RETURNING id ", genre.Name, genre.Description, genre.CreatedAt))

	for rows.Next() {

		err := rows.Scan(&genre.ID)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("Gênero Inserido: ID=%s Name=%s Description=%s MovieCount=%d CreatedAt=%s", genre.ID, genre.Name, genre.Description, genre.MovieCount, genre.CreatedAt)

	json.NewEncoder(w).Encode(&genre)
}
