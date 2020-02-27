package characters

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/kjintroverted/wizardsBrew/data/tasks"

	"github.com/kjintroverted/wizardsBrew/api/backgrounds"
	"github.com/kjintroverted/wizardsBrew/api/classes"
	"github.com/kjintroverted/wizardsBrew/api/feats"
	"github.com/kjintroverted/wizardsBrew/api/items"
	"github.com/kjintroverted/wizardsBrew/api/races"
	"github.com/kjintroverted/wizardsBrew/api/spells"

	"github.com/gorilla/mux"
	"github.com/kjintroverted/wizardsBrew/psql"
)

// UpsertPC Create a new PC if no ID is given
// or updates a PC for the given ID
func UpsertPC(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]interface{})
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()

	service := NewPCService(NewPCRepo(db))

	b, _ := ioutil.ReadAll(r.Body)
	var data PC
	if err = json.Unmarshal(b, &data); err != nil {
		fmt.Println("ERR", err)
	}

	uid := r.Context().Value("uid").(string)
	if data.ID != "" {
		if auth := service.Authorized(fmt.Sprintf("%v", data.ID), uid); !auth {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized UID: " + uid))
			return
		}
	}

	id, err := service.Upsert(data, uid)
	if err == nil {
		res["id"] = id
		b, _ = json.Marshal(res)
		w.Write(b)
		return
	}

	b, _ = json.Marshal(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(b)
}

// RequestAccess will post a uid to a characters auth requests
func RequestAccess(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()

	service := NewPCService(NewPCRepo(db))

	id := mux.Vars(r)["id"]
	uid := r.Context().Value("uid").(string)

	if err := service.RequestAuth(id, uid); err != nil {
		b, _ := json.Marshal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b)
		return
	}

	w.Write([]byte("Request sent."))
}

// PartyInvite will post a party id to a characters party invites
func PartyInvite(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()

	service := NewPCService(NewPCRepo(db))

	id := mux.Vars(r)["id"]
	party := r.URL.Query().Get("party")

	if err := service.Invite(id, party); err != nil {
		b, _ := json.Marshal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b)
		return
	}

	w.Write([]byte("Invite sent."))
}

// DeletePC deletes a PC
func DeletePC(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()

	service := NewPCService(NewPCRepo(db))

	// LOAD PATH PARAMS
	pathParams := mux.Vars(r)
	// GET AUTH UID
	uid := r.Context().Value("uid").(string)

	if id, ok := pathParams["id"]; ok {
		if !service.Authorized(id, uid) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		if err = service.Delete(id); err == nil {
			w.Write([]byte("Deleted"))
			return
		}
	}

	b, _ := json.Marshal(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(b)
}

// PlayableCharacters handles all GET req for PCs
func PlayableCharacters(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()
	service := NewPCService(NewPCRepo(db))

	var res []byte

	// LOAD PATH PARAMS
	pathParams := mux.Vars(r)
	detail := r.URL.Query().Get("detail") == "true"

	// GET AUTH UID
	uid := r.Context().Value("uid").(string)

	if id, ok := pathParams["id"]; ok { // GET ONE BY ID
		var pc *PC
		if pc, err = service.FindByID(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if !detail {
			res, _ = json.Marshal(pc)
			w.Write(res)
			return
		}

		data := getCharacterData(*pc, uid, service, db)
		res, _ = json.Marshal(data)
		w.Write(res)

	} else {
		var data []map[string]interface{}
		var characters []PC

		if q := r.URL.Query().Get("name"); q != "" {
			if characters, err = service.Search(q); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		} else {
			if characters, err = service.FindByUser(uid); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		}

		if !detail {
			res, _ = json.Marshal(characters)
			w.Write(res)
			return
		}

		for _, pc := range characters {
			data = append(data, getCharacterData(pc, uid, service, db))
		}
		res, _ = json.Marshal(data)
		w.Write(res)
	}
}

func getCharacterData(pc PC, uid string, service PCService, db *sql.DB) map[string]interface{} {
	var err error
	data := make(map[string]interface{})

	data["info"] = pc

	lvl := tasks.GetLevelInfo(pc.XP)
	data["level"] = lvl

	data["authorized"] = service.AuthorizedLocal(pc, uid)

	var background *backgrounds.Background
	var class *classes.Class
	if race, err := races.NewRaceService(races.NewRaceRepo(db)).FindByID(strconv.Itoa(pc.RaceID)); err == nil {
		data["race"] = race
	}
	if class, err = classes.NewClassService(classes.NewClassRepo(db)).FindByID(strconv.Itoa(pc.ClassID)); err == nil {
		data["class"] = class
	}
	if background, err = backgrounds.NewBackgroundService(backgrounds.NewBackgroundRepo(db)).FindByID(strconv.Itoa(pc.BackgroundID)); err == nil {
		data["background"] = background
	}
	if equipment, err := items.NewItemService(items.NewItemRepo(db)).FindByIDs(pc.EquipmentIDs); err == nil {
		data["equipment"] = equipment
	}
	if weapons, err := items.NewItemService(items.NewItemRepo(db)).FindByIDs(pc.WeaponIDs); err == nil {
		data["weapons"] = weapons
	}

	var inventoryIDs []psql.NullInt
	for _, i := range pc.Inventory {
		inventoryIDs = append(inventoryIDs, psql.NullInt{sql.NullInt64{int64(i.ID), true}})
	}
	if inventory, err := items.NewItemService(items.NewItemRepo(db)).FindByIDs(inventoryIDs); err == nil {
		data["inventory"] = inventory
	}
	if spells, err := spells.NewSpellService(spells.NewSpellRepo(db)).FindByIDs(pc.SpellIDs); err == nil {
		data["spells"] = spells
	}

	var featArr []feats.Feat
	featService := feats.NewFeatService(feats.NewFeatRepo(db))
	if f, err := featService.FindByIDs(pc.SpecFeatIDs); err == nil {
		featArr = append(featArr, f...)
	}
	var opts = map[string][]string{
		"class":    []string{class.Name},
		"subclass": []string{pc.Subclass},
		"level":    []string{strconv.Itoa(lvl.Level)},
	}
	if f, err := featService.List(opts); err == nil {
		featArr = append(featArr, f...)
	}
	opts = map[string][]string{
		"background": []string{background.Name},
	}
	if f, err := featService.List(opts); err == nil {
		featArr = append(featArr, f...)
	}

	data["features"] = featArr

	return data
}
