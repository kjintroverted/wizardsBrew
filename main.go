package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kjintroverted/wizardsBrew/api/characters"

	"github.com/kjintroverted/wizardsBrew/api"

	"github.com/gorilla/mux"
	"github.com/kjintroverted/wizardsBrew/quiz"

	"github.com/kjintroverted/wizardsBrew/data/tasks"
)

func main() {
	var command string

	if len(os.Args) == 1 {
		goto API
	}

	command = os.Args[1]
	switch command {
	case "sql-gen":
		tasks.GenerateClassInserts()
		tasks.GenerateItemInserts()
		tasks.GenerateSpellInserts()
		tasks.GenerateFeatInserts()
		tasks.GenerateBackgroundInserts()
	case "sql":
		tasks.GenerateClassInserts()
		tasks.GenerateItemInserts()
		tasks.GenerateSpellInserts()
		tasks.GenerateFeatInserts()
		tasks.GenerateBackgroundInserts()
		tasks.ExecSQL("data")
	case "quiz":
		quiz.BeginQuiz()
	}
	return

API:
	// SET PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// GET MULTIPLEX
	mux := createMux()

	// START SERVER
	fmt.Println("Server listening on", port)
	if error := http.ListenAndServe(":"+port, mux); error != nil {
		fmt.Println("ERROR", error)
	}
}

// CREATE MULTIPLEX PATH HANDLER
func createMux() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api", root)

	// SRD ENDPOINTS
	r.HandleFunc("/api/items", api.Items).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/items/{id}", api.Items).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/spells", api.Spells).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/spells/{id}", api.Spells).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/races", api.Races).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/races/{id}", api.Races).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/classes", api.Classes).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/classes/{id}", api.Classes).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/feats", api.Feats).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/feats/{id}", api.Feats).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/bg", api.Backgrounds).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/bg/{id}", api.Backgrounds).Methods("GET", "OPTIONS")

	// PLAYABLE CHARACTERS
	r.HandleFunc("/api/data/pc/{id}", characters.PlayableCharacters).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/data/pc", characters.PlayableCharacters).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/data/pc", characters.UpsertPC).Methods("POST", "PUT", "OPTIONS")
	r.HandleFunc("/api/data/pc/{id}", characters.DeletePC).Methods("DELETE", "OPTIONS")

	// USERS
	r.HandleFunc("/api/user", getUserData).Methods("GET", "OPTIONS")

	r.Use(cors)
	r.Use(userAuth)

	return r
}

// ROOT GETTER
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wizard's Brew API")
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Enabling CORS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if strings.Contains(r.URL.Path, "/api/data") {
			w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
		}
		if r.Method == "OPTIONS" {
			w.Header().Add("Connection", "keep-alive")
			w.Header().Add("Access-Control-Max-Age", "86400")
			w.Write([]byte(""))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
