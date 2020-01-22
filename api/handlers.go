package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kjintroverted/wizardsBrew/api/classes"
	"github.com/kjintroverted/wizardsBrew/api/feats"
	"github.com/kjintroverted/wizardsBrew/api/items"
	"github.com/kjintroverted/wizardsBrew/api/races"
	"github.com/kjintroverted/wizardsBrew/api/spells"
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

	var res []byte

	// LOAD PATH PARAMS
	pathParams := mux.Vars(r)

	if id, ok := pathParams["id"]; ok { // GET ONE BY ID
		var item *items.Item
		if item, err = service.FindByID(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		res, _ = json.Marshal(item)
		w.Write(res)
		return
	}

	// FIND BY TYPE
	itemType := r.URL.Query().Get("type")
	var items []items.Item
	switch itemType {
	case "": // GEAR AND ITEMS
		if items, err = service.FindItems(); err != nil {
			goto Fail
		}
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
	res, _ = json.Marshal(items)
	w.Write(res)
	return

Fail:
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

// Spells handles all Spell req
func Spells(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()
	service := spells.NewSpellService(spells.NewSpellRepo(db))

	var res []byte

	// LOAD PATH PARAMS
	pathParams := mux.Vars(r)

	if id, ok := pathParams["id"]; ok { // GET ONE BY ID
		var spell *spells.Spell
		if spell, err = service.FindByID(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		res, _ = json.Marshal(spell)
		w.Write(res)
		return
	}

	var spells []spells.Spell
	if spells, err = service.List(r.URL.Query()); err == nil {
		res, _ = json.Marshal(spells)
		w.Write(res)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

// Races handles all Race req
func Races(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()
	service := races.NewRaceService(races.NewRaceRepo(db))

	var res []byte

	// LOAD PATH PARAMS
	pathParams := mux.Vars(r)

	if id, ok := pathParams["id"]; ok { // GET ONE BY ID
		var race *races.Race
		if race, err = service.FindByID(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		res, _ = json.Marshal(race)
		w.Write(res)
		return
	}
	if races, err := service.List(); err == nil {
		res, _ = json.Marshal(races)
		w.Write(res)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

// Classes handles all Class req
func Classes(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()
	service := classes.NewClassService(classes.NewClassRepo(db))

	var res []byte

	// LOAD PATH PARAMS
	pathParams := mux.Vars(r)

	if id, ok := pathParams["id"]; ok { // GET ONE BY ID
		var class *classes.Class
		if class, err = service.FindByID(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		res, _ = json.Marshal(class)
		w.Write(res)
		return
	}
	var classes []classes.Class
	if classes, err = service.List(); err == nil {
		res, _ = json.Marshal(classes)
		w.Write(res)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

// Feats handles all Feat req
func Feats(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()
	service := feats.NewFeatService(feats.NewFeatRepo(db))

	var res []byte

	// LOAD PATH PARAMS
	pathParams := mux.Vars(r)

	if id, ok := pathParams["id"]; ok { // GET ONE BY ID
		var feat *feats.Feat
		if feat, err = service.FindByID(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		res, _ = json.Marshal(feat)
		w.Write(res)
		return
	}

	var feats []feats.Feat
	if feats, err = service.List(r.URL.Query()); err == nil {
		res, _ = json.Marshal(feats)
		w.Write(res)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
