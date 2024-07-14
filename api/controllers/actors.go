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

func GetActors(w http.ResponseWriter, r *http.Request) {
	var actors []models.Actor
	rows := db.Query("SELECT * FROM actor")

	for rows.Next() {
		var actor models.Actor
		err := rows.StructScan(&actor)
		if err != nil {
			log.Fatalln(err)
		}
		actors = append(actors, actor)
	}

	utils.ConsoleLog(w, r)
	log.Printf("%d Registros retornados", len(actors))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actors)
}

func GetActorsByID(w http.ResponseWriter, r *http.Request) {
	utils.ConsoleLog(w, r)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var actor models.Actor
	rows := db.Query("SELECT * FROM actor WHERE id = " + params["id"])
	for rows.Next() {
		err := rows.StructScan(&actor)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if actor.ID == "" {
		NotFound(w, r, fmt.Sprintf("Nenhum registro para ID=%s encontrado", params["id"]))
		log.Printf("Nenhum registro para ID=%s encontrado", params["id"])
		return
	}

	json.NewEncoder(w).Encode(&actor)
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)

	var actor models.Actor

	err := json.Unmarshal(data, &actor)
	utils.ConsoleLog(w, r)

	if err != nil {
		log.Fatalf("Erro ao converter JSON para struct: %s", err)
	}

	rows := db.Query(fmt.Sprintf("INSERT INTO actor (name, birth_year, nationality, movies_starred) VALUES ('%s',%d,'%s',%d) RETURNING id", actor.Name, actor.BirthYear, actor.Nationality, actor.MoviesStarred))

	for rows.Next() {

		err := rows.Scan(&actor.ID)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("Ator Inserido: ID=%s Name=%s BirthYear=%d Nationality=%s MoviesStarred=%d", actor.ID, actor.Name, actor.BirthYear, actor.Nationality, actor.MoviesStarred)

	json.NewEncoder(w).Encode(&actor)
}
