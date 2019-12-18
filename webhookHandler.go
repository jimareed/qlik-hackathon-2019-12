package main

import (
	"log"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {

	event := r.Header.Get("x-github-event")

	log.Print("Github Event: ")
	log.Print(event)

	apiKey := ""
	tenantUrl := ""
	
	items, err := getItem(apiKey, tenantUrl, "Bugs")
	if err != nil {
		log.Print("Get Item error") 
	} else {
		for _, i := range items.Data {
			log.Print("reloading " + i.Name)
			_, err := reload(apiKey, tenantUrl, i.ResourceId)
			if err != nil {
				log.Print("Reload error") 
			}
		}
	}


	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

