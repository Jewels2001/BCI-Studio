package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"backend/server/routes"
	"backend/server/types"
	"backend/server/util"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	PORT = 8080
)

func greet(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: /")
	util.RespondWithJSON(w, http.StatusAccepted, &types.Message{Message: fmt.Sprintf("Hello World! %s", time.Now())})
}

func main() {
	// Setup routes
	r := mux.NewRouter()
	r.HandleFunc("/", greet)
	r.HandleFunc("/upload", routes.Upload).Methods("POST")
	r.HandleFunc("/files", routes.GetFiles).Methods("GET")
	r.HandleFunc("/download/{filename}", routes.Download).Methods("GET")

	log.Printf("Now serving on %d...\n", PORT)

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), corsWrapper.Handler(r))
}
