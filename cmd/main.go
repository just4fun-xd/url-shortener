package main

import (
	"log"
	"net/http"
	"url-shortener/pkg/api"
	"url-shortener/pkg/storage"
)

func main() {
	db, err := storage.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	api.DB = db

	router := api.SetupRouter()
	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
