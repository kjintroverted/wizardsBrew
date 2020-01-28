package main

import (
	"fmt"
	"net/http"
	"os"

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
		port = "4000"
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
	r.HandleFunc("/api/items", api.Items)
	r.HandleFunc("/api/items/{id}", api.Items)
	r.HandleFunc("/api/spells", api.Spells)
	r.HandleFunc("/api/spells/{id}", api.Spells)
	r.HandleFunc("/api/races", api.Races)
	r.HandleFunc("/api/races/{id}", api.Races)
	r.HandleFunc("/api/classes", api.Classes)
	r.HandleFunc("/api/classes/{id}", api.Classes)
	r.HandleFunc("/api/feats", api.Feats)
	r.HandleFunc("/api/feats/{id}", api.Feats)
	r.HandleFunc("/api/back", api.Backgrounds)
	r.HandleFunc("/api/back/{id}", api.Backgrounds)
	r.HandleFunc("/api/pc", characters.UpsertPC).Methods("POST", "PUT")

	r.Use(enableCORS)

	return r
}

// ROOT GETTER
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wizard's Brew API")
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
