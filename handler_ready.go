package main

import (
	"log"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	log.Println("In handlerReadiness: Sending Status OK")
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
