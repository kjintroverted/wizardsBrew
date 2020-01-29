package characters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

	id, err := service.Upsert(data)
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

	if id, ok := pathParams["id"]; ok { // GET ONE BY ID
		var pc *PC
		if pc, err = service.FindByID(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		res, _ = json.Marshal(pc)
		w.Write(res)
		return
	}

	if err == nil {
		err = fmt.Errorf("no logic found for request")
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
