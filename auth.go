package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
)

func getUserData(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		fmt.Println("APP ERROR:", err.Error())
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	uid := r.URL.Query().Get("uid")
	if user, err := client.GetUser(ctx, uid); err == nil {
		b, _ := json.Marshal(map[string]interface{}{
			"uid":      uid,
			"name":     user.UserInfo.DisplayName,
			"email":    user.Email,
			"photoURL": user.PhotoURL,
		})
		w.Write(b)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func userAuth(next http.Handler) http.Handler {
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
		if !strings.Contains(r.URL.Path, "/api/data") && r.Method == "GET" {
			next.ServeHTTP(w, r)
			return
		}
		auth := r.Header.Get("Authorization")
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Error completing request:\n %+v", err)))
			}
		}()

		mode := strings.Split(auth, " ")[0]
		token := strings.Split(auth, " ")[1]
		var uid string

		if mode == "Bearer" {
			authToken, err := client.VerifyIDToken(ctx, token)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(err.Error()))
				return
			}
			uid = authToken.UID
		} else if mode == "dev" && os.Getenv("ENV") != "PROD" {
			user, err := client.GetUserByEmail(ctx, token)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(err.Error()))
				return
			}
			uid = user.UserInfo.UID
		}

		// fmt.Println(authToken.UID)
		ctx := context.WithValue(r.Context(), "uid", uid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
