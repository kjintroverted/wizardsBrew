package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
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
	r.HandleFunc("/api/items", api.Items).Methods("GET")
	r.HandleFunc("/api/items/{id}", api.Items).Methods("GET")
	r.HandleFunc("/api/spells", api.Spells).Methods("GET")
	r.HandleFunc("/api/spells/{id}", api.Spells).Methods("GET")
	r.HandleFunc("/api/races", api.Races).Methods("GET")
	r.HandleFunc("/api/races/{id}", api.Races).Methods("GET")
	r.HandleFunc("/api/classes", api.Classes).Methods("GET")
	r.HandleFunc("/api/classes/{id}", api.Classes).Methods("GET")
	r.HandleFunc("/api/feats", api.Feats).Methods("GET")
	r.HandleFunc("/api/feats/{id}", api.Feats).Methods("GET")
	r.HandleFunc("/api/bg", api.Backgrounds).Methods("GET")
	r.HandleFunc("/api/bg/{id}", api.Backgrounds).Methods("GET")

	// PLAYABLE CHARACTERS
	r.Handle("/api/pc/{id}", googleUser(http.HandlerFunc(characters.PlayableCharacters))).Methods("GET")
	r.Handle("/api/pc", googleUser(http.HandlerFunc(characters.PlayableCharacters))).Methods("GET")
	r.Handle("/api/pc", googleUser(http.HandlerFunc(characters.UpsertPC))).Methods("POST", "PUT")
	r.Handle("/api/pc/{id}", googleUser(http.HandlerFunc(characters.DeletePC))).Methods("DELETE")

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

func googleUser(next http.Handler) http.Handler {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(fmt.Sprintf("Error completing request:\n %v", err)))
			}
		}()

		token := strings.Split(auth, " ")[1]

		authToken, err := client.VerifyIDToken(ctx, token)
		// fmt.Println(authToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), "uid", authToken.UID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
