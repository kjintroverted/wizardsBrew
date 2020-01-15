package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kjintroverted/wizardsBrew/api/items"
	"github.com/kjintroverted/wizardsBrew/psql"

	"github.com/gorilla/mux"
)

// Items handles all Item req
func Items(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()
	service := items.NewItemService(items.NewItemRepo(db))

	// LOAD PATH PARAMS
	pathParams := mux.Vars(r)

	if id, ok := pathParams["id"]; ok { // GET ONE
		item, e := service.FindByID(id)
		if e != nil {
			err = e
			goto Fail
		}
		res, _ := json.Marshal(item)
		w.Write(res)
		return
	}

Fail:
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
