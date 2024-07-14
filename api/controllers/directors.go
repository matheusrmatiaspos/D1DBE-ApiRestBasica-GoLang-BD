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

func GetDirectors(w http.ResponseWriter, r *http.Request) {
	var directors []models.Director
	rows := db.Query("SELECT * FROM director")

	for rows.Next() {
		var director models.Director
		err := rows.StructScan(&director)
		if err != nil {
			log.Fatalln(err)
		}
		directors = append(directors, director)
	}

	utils.ConsoleLog(w, r)
	log.Printf("%d Registros retornados", len(directors))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(directors)
}

func GetDirectorById(w http.ResponseWriter, r *http.Request) {
	utils.ConsoleLog(w, r)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var director models.Director
	rows := db.Query("SELECT * FROM director WHERE id = " + params["id"])
	for rows.Next() {
		err := rows.StructScan(&director)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if director.ID == "" {
		NotFound(w, r, fmt.Sprintf("Nenhum registro para ID=%s encontrado", params["id"]))
		log.Printf("Nenhum registro para ID=%s encontrado", params["id"])
		return
	}

	json.NewEncoder(w).Encode(&director)
}

func CreateDirector(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)

	var director models.Director

	err := json.Unmarshal(data, &director)
	utils.ConsoleLog(w, r)

	if err != nil {
		log.Fatalf("Erro ao converter JSON para struct: %s", err)
	}

	rows := db.Query(fmt.Sprintf("INSERT INTO director(name,birth_year,nationality,movies_directed) VALUES ('%s',%d,'%s',%d) RETURNING id", director.Name, director.BirthYear, director.Nationality, director.MoviesDirected))

	for rows.Next() {

		err := rows.Scan(&director.ID)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("Diretor Inserido: ID=%s Name=%s BirthYear=%d Nationality=%s MoviesDirected=%d", director.ID, director.Name, director.BirthYear, director.Nationality, director.MoviesDirected)

	json.NewEncoder(w).Encode(&director)
}
