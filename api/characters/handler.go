package characters

import (
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
	if data.ID != 0 {
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

	if id, ok := pathParams["id"]; ok {
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
	// GET AUTH UID
	uid := r.Context().Value("uid").(string)

	if id, ok := pathParams["id"]; ok { // GET ONE BY ID
		data := make(map[string]interface{})

		var pc *PC
		if pc, err = service.FindByID(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		data["info"] = pc

		lvl := tasks.GetLevelInfo(pc.XP)
		data["level"] = lvl

		data["authorized"] = service.AuthorizedLocal(*pc, uid)

		var race *races.Race
		var background *backgrounds.Background
		var class *classes.Class
		if race, err = races.NewRaceService(races.NewRaceRepo(db)).FindByID(strconv.Itoa(pc.RaceID)); err == nil {
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
		if inventory, err := items.NewItemService(items.NewItemRepo(db)).FindByIDs(pc.InventoryIDs); err == nil {
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

		res, _ = json.Marshal(data)
		w.Write(res)
		return
	}

	if err == nil {
		err = fmt.Errorf("no logic found for request")
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
