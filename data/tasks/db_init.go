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

// HELPERS

type section struct {
	Title string
	Body  []interface{}
}

// san will sanitize string for psql statements
func san(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}

// san will sanitize string for psql statements
func sanAll(s []interface{}) (res []interface{}) {
	for _, inter := range s {
		if _, ok := inter.(string); !ok {
			continue
		}
		res = append(res, strings.ReplaceAll(inter.(string), "'", "''"))
	}
	return
}

func stripFilter(s string) string {
	if !strings.Contains(s, "{@filter") {
		return s
	}

	str := string(s[9:])
	return strings.Split(str, "|")[0]
}

func simpleStrArray(arr []interface{}) string {
	if len(arr) == 0 {
		return "null"
	}
	s := "array["
	for _, x := range arr {
		if str, ok := x.(string); ok {
			s += fmt.Sprintf("'%s',", san(str))
		}
	}
	return strings.Trim(s, ",") + "]"
}

func simpleArray(arr []interface{}) string {
	if len(arr) == 0 {
		return "null"
	}
	s := "array["
	for _, x := range arr {
		s += fmt.Sprintf("%s,", x)
	}
	return strings.Trim(s, ",") + "]"
}

func rowString(t string, arr ...interface{}) string {
	if len(arr) == 0 {
		return "null"
	}
	s := "row("
	for _, x := range arr {
		s += fmt.Sprintf("%s,", x)
	}
	return fmt.Sprintf("%s)::%s", strings.Trim(s, ","), t)
}

func join(arr []interface{}, sep string) (j string) {
	for _, s := range arr {
		j += fmt.Sprintf("%s%s", s, sep)
	}
	j = strings.Trim(j, sep)
	return
}
