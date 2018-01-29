package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	user     = "goform"
	host     = "mysql"
	database = "goform"
	table    = "postdata"
	tpl      = "/usr/local/goform_sql/html/index.html.tpl"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Top(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	value := r.FormValue("key")
	if len(value) != 0 {
		Insert(value)
	}

	list := Select()
	t := template.Must(template.ParseFiles(tpl))
	err := t.ExecuteTemplate(w, "index.html.tpl", list)
	CheckError(err)

}

func Connect() string {
	password := os.Getenv("USER_PASSWORD")
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", user, password, host, database)
	return dbConnection
}

func Insert(w string) {
	db, err := sql.Open("mysql", Connect())
	CheckError(err)
	defer db.Close()

	data := fmt.Sprintf("INSERT INTO %s(val) values('%s')", table, w)
	_, err = db.Exec(data)
	CheckError(err)
}

func Select() []string {
	db, err := sql.Open("mysql", Connect())
	CheckError(err)
	defer db.Close()

	rows, err := db.Query("SELECT id, val, created FROM " + table)
	CheckError(err)
	defer rows.Close()

	var s []string
	for rows.Next() {
		var id int
		var val string
		var created string

		err := rows.Scan(&id, &val, &created)
		CheckError(err)

		result := fmt.Sprintf("%d | %s | %s", id, created, val)
		s = append(s, result)
	}
	return s
}

func main() {
	http.HandleFunc("/", Top)

	fmt.Printf("Starting web server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
