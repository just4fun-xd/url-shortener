package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"url-shortener/pkg/storage"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/shorten", CreateShortURL).Methods("POST")
	router.HandleFunc("/{short_url}", RedirectURL).Methods("GET")
	return router
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the URL Shortener API!")
	fmt.Fprintln(w, "Use POST /shorten to create a short URL.")
	fmt.Fprintln(w, "Use GET /{short_url} to redirect to the original URL.")
}

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		OriginalURL string `json:"original_url"`
	}
	json.NewDecoder(r.Body).Decode(&payload)

	shortURL := generateShortURL()
	newURL := storage.URL{
		ShortURL:    shortURL,
		OriginalURL: payload.OriginalURL,
	}

	DB.Create(&newURL)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newURL)
}

func generateShortURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["short_url"]

	var url storage.URL
	if err := DB.First(&url, "short_url = ?", shortURL).Error; err != nil {
		http.Error(w, "Ooops URL not found", http.StatusNotFound)
		return
	}

	DB.Model(&url).Update("visit_count", url.VisionCount+1)
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
