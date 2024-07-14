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

func GetMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	rows := db.Query("SELECT m.id, m.title, d.name AS director, g.description AS genre, m.release_year  FROM movie m INNER JOIN director d ON m.director=d.id INNER JOIN genre g ON m.genre = g.id")

	for rows.Next() {
		var movie models.Movie
		err := rows.StructScan(&movie)
		if err != nil {
			log.Fatalln(err)
		}
		movies = append(movies, movie)
	}

	utils.ConsoleLog(w, r)
	log.Printf("%d Registros retornados", len(movies))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	utils.ConsoleLog(w, r)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var movie models.Movie
	rows := db.Query(fmt.Sprintf("SELECT m.id, m.title, d.name AS director, g.name AS genre, m.release_year  FROM movie m INNER JOIN director d ON m.director=d.id INNER JOIN genre g ON m.genre = g.id WHERE m.id=%s", params["id"]))
	for rows.Next() {
		err := rows.StructScan(&movie)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if movie.ID == "" {
		NotFound(w, r, fmt.Sprintf("Nenhum registro para ID=%s encontrado", params["id"]))
		log.Printf("Nenhum registro para ID=%s encontrado", params["id"])
		return
	}

	json.NewEncoder(w).Encode(&movie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)

	var movie models.Movie

	err := json.Unmarshal(data, &movie)
	utils.ConsoleLog(w, r)

	if err != nil {
		log.Fatalf("Erro ao converter JSON para struct: %s", err)
	}

	rows := db.Query(fmt.Sprintf("INSERT INTO movie (title,director,genre,release_year) VALUES ('%s',%s,%s,%d) RETURNING id", movie.Title, movie.Director, movie.Genre, movie.ReleaseYear))

	for rows.Next() {

		err := rows.Scan(&movie.ID)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("Filme Inserido: ID=%s Title=%s DirectorID=%s GenreID=%s ReleaseYear=%d", movie.ID, movie.Title, movie.Director, movie.Genre, movie.ReleaseYear)

	json.NewEncoder(w).Encode(&movie)
}
