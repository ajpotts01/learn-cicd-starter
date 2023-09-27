package main

import (
	"log"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	log.Println("In handlerReadiness: Sending Status Already Reported")
	respondWithJSON(w, http.StatusAlreadyReported, map[string]string{"status": "ok"})
}
