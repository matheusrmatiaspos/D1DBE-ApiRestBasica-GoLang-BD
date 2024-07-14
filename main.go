package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang-BD/server"
)

func main() {
	router := mux.NewRouter()
	server.Start(router, 8080)
	log.Println("Sevidor rodando em: http://localhost:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
