package parties

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjintroverted/wizardsBrew/psql"
)

// UpsertParty Create a new Party if no ID is given
// or updates a Party for the given ID
func UpsertParty(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]interface{})
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()

	service := NewPartyService(NewPartyRepo(db))

	b, _ := ioutil.ReadAll(r.Body)
	var data Party
	if err = json.Unmarshal(b, &data); err != nil {
		fmt.Println("ERR", err)
	}

	uid := r.Context().Value("uid").(string)

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

// JoinParty adds a member to the specified party
func JoinParty(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()

	service := NewPartyService(NewPartyRepo(db))

	id := mux.Vars(r)["id"]
	member := r.URL.Query().Get("member")

	if err := service.Join(id, member); err != nil {
		b, _ := json.Marshal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b)
		return
	}

	w.Write([]byte("Welcome to the Party"))
}

// KickMember removes a member from the specified party
func KickMember(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()

	service := NewPartyService(NewPartyRepo(db))

	id := mux.Vars(r)["id"]
	member := r.URL.Query().Get("member")
	uid := r.Context().Value("uid").(string)

	if err := service.KickMember(id, uid, member); err != nil {
		b, _ := json.Marshal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b)
		return
	}

	w.Write([]byte("Bye"))
}

// DeleteParty deletes a Party
func DeleteParty(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()

	service := NewPartyService(NewPartyRepo(db))

	// LOAD PATH PARAMS
	pathParams := mux.Vars(r)
	// GET AUTH UID
	uid := r.Context().Value("uid").(string)

	if id, ok := pathParams["id"]; ok {
		if err = service.Delete(id, uid); err == nil {
			w.Write([]byte("Deleted"))
			return
		}
	}

	b, _ := json.Marshal(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(b)
}

// Parties handles all GET req for Parties
func Parties(w http.ResponseWriter, r *http.Request) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()
	service := NewPartyService(NewPartyRepo(db))

	var res []byte

	// LOAD PATH PARAMS
	pathParams := mux.Vars(r)
	// GET AUTH UID
	uid := r.Context().Value("uid").(string)

	if id, ok := pathParams["id"]; ok { // GET ONE BY ID
		var party *Party
		if party, err = service.FindByID(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		res, _ = json.Marshal(party)
		w.Write(res)
	} else {
		id = r.URL.Query().Get("member")
		if id == "" {
			id = uid
		}

		if parties, err := service.FindByMember(id); err == nil {
			res, _ = json.Marshal(parties)
			w.Write(res)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
