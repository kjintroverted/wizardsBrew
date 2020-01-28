package characters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
