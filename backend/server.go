package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"backend/server/types"
	"backend/server/util"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	PORT = 8080
)

func greet(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: greet")
	util.RespondWithJSON(w, 200, fmt.Sprintf("Hello World! %s", time.Now()))
}

func uploadCSV(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: uploadCSV")
	var H types.History
	err := json.NewDecoder(r.Body).Decode(&H)
	if err != nil {
		log.Println(err.Error())
		util.RespondWithError(w, http.StatusBadRequest, "invalid fields")
		return
	}
	util.DisplayH(H)
	util.RespondWithJSON(w, http.StatusAccepted, fmt.Sprintf("Data Recieved %s", time.Now()))
}

func main() {
	// Setup routes
	r := mux.NewRouter()
	r.HandleFunc("/", greet)
	r.HandleFunc("/upload", uploadCSV).Methods("POST")

	log.Printf("Now serving on %d...\n", PORT)

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), corsWrapper.Handler(r))
}
