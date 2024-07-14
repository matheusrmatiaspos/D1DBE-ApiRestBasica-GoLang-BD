package controllers

import (
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request, message string){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, message)
}