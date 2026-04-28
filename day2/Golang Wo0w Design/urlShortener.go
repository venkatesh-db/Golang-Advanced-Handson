package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const shortLength = 6

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=pass dbname=shortener sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
}

func generateShortCode() string {
	b := make([]byte, shortLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

type ShortenRequest struct {
	URL string `json:"url"`
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	code := generateShortCode()
	_, err := db.Exec("INSERT INTO url_mappings (short_code, original_url) VALUES ($1, $2)", code, req.URL)
	if err != nil {
		http.Error(w, "Insert failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"short_url": fmt.Sprintf("http://localhost:8080/%s", code),
	})
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	var original string
	err := db.QueryRow("SELECT original_url FROM url_mappings WHERE short_code = $1", code).Scan(&original)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	db.Exec("UPDATE url_mappings SET click_count = click_count + 1 WHERE short_code = $1", code)
	http.Redirect(w, r, original, http.StatusFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", shortenHandler).Methods("POST")
	r.HandleFunc("/{code}", redirectHandler).Methods("GET")

	log.Println("🚀 Server on :9080")
	log.Fatal(http.ListenAndServe(":9080", r))
}
