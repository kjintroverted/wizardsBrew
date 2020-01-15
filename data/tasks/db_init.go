package tasks

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kjintroverted/wizardsBrew/psql"
)

// ExecSQL runs all .sql files
func ExecSQL(dir string) {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()

	root, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("ERR", err)
	}

	execAll(root, dir, db)
}

func execAll(files []os.FileInfo, root string, db *sql.DB) {
	for _, info := range files {
		if info.IsDir() { // RECUR CALL ON CHILDREN
			children, _ := ioutil.ReadDir(root + "/" + info.Name())
			execAll(children, root+"/"+info.Name(), db)
			continue
		}

		if !strings.Contains(info.Name(), "sql") { // SKIP NON_SQL FILES
			continue
		}

		raw, err := ioutil.ReadFile(root + "/" + info.Name())
		if err != nil {
			fmt.Println("ERR", err)
			return
		}

		fmt.Println("Running " + root + "/" + info.Name())
		parseAndExec(string(raw), db)
	}
}

func parseAndExec(queryBlob string, db *sql.DB) {
	qArr := strings.Split(queryBlob, ";\n")
	rows := 0
	for _, q := range qArr {
		q = strings.Trim(q, " \n") + ";"
		res, err := db.Exec(q)
		if err != nil {
			fmt.Println("ERR", err)
			continue
		}
		r, _ := res.RowsAffected()
		rows += int(r)
	}
	fmt.Printf("\t%d rows affected\n", rows)
}
