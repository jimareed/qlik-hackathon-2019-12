package main

import (
	"log"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {

	event := r.Header.Get("x-github-event")

	log.Print("Github Event: ")
	log.Print(event)
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

