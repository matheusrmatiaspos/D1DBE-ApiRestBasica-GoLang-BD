package utils

import (
	"log"
	"net/http"
)

func ConsoleLog(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s%s [%s]",r.Host,r.URL,r.Method)
}