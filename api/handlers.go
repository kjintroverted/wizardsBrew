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

	if id, ok := pathParams["id"]; ok { // GET ONE BY ID
		var item *items.Item
		if item, err = service.FindByID(id); err != nil {
			goto Fail
		}
		res, _ := json.Marshal(item)
		w.Write(res)
		return
	} else if itemType := r.URL.Query().Get("type"); itemType != "" { // FIND BY TYPE
		var items []items.Item
		switch itemType {
		case "weapon": // WEAPONS
			if items, err = service.FindWeapons(); err != nil {
				goto Fail
			}
		case "armor": // ARMOR
			if items, err = service.FindArmor(); err != nil {
				goto Fail
			}
		default:
			err = fmt.Errorf("Could not match type: %s", itemType)
			goto Fail
		}
		res, _ := json.Marshal(items)
		w.Write(res)
		return
	}
	err = fmt.Errorf("Requested endpoint not found")

Fail:
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
