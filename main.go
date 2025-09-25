package main

import (
	"log"
	"net/http"
	//"os"
	"time"

	"github.com/firemanm/go_crud/handlers"

	"github.com/firemanm/go_crud/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env for local development
	godotenv.Load()

	// DB init
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// table create
	err = database.CreateTable(db)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	// handlers init
	userHandler := handlers.NewUserHandler(db)

	// mux routes init
	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	log.Printf("Routes registered...")

	// health check handler
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")
	log.Printf("Health route registered...")


	// run server
	
	port := "8080"
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(server.ListenAndServe())
}
