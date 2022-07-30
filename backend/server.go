package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	// Setup routes
	r := mux.NewRouter()
	r.HandleFunc("/", greet)
	// r.HandleFunc("/register", routes.CreateUser).Methods("POST")
	// r.HandleFunc("/login", routes.Login).Methods("POST")
	r.HandleFunc("/upload", uploadCSV).Methods("POST")

	// Routes with Auth middleware
	// authRoute := r.Methods("GET", "POST").Subrouter()
	// authRoute.HandleFunc("/users", routes.GetAllUsers).Methods("GET")
	// authRoute.HandleFunc("/users/{id}", routes.GetUser).Methods("GET")
	// authRoute.HandleFunc("/days", routes.GetDays).Methods("GET")
	// authRoute.HandleFunc("/days", routes.CreateDay).Methods("POST")
	// authRoute.HandleFunc("/logout", routes.Logout)
	// authRoute.Use(auth.AuthMiddleware)

	fmt.Println("Now serving on 8080...")

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	http.ListenAndServe(":8080", corsWrapper.Handler(r))
}
