package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/proullon/ramsql/driver"
)

func initDB() *sql.DB {
	db, err := sql.Open("ramsql", "TestLoadUserAddresses")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("CREATE TABLE user (id BIGSERIAL PRIMARY KEY, username TEXT, password TEXT);")
	db.Exec("INSERT INTO user (username, password) VALUES ('admin', 'root');")
	db.Exec("INSERT INTO user (username, password) VALUES ('sulaev', '123123');")
	db.Exec("INSERT INTO user (username, password) VALUES ('k.kitsuragi', 'revachol');")

	return db
}

func main() {
	db := initDB()

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")

		// SELECT id FROM user WHERE username=a.sulaev AND password=123
		query := "SELECT id FROM user WHERE username=" + username + " AND password=" + password

		row := db.QueryRow(query)
		var id int
		if err := row.Scan(&id); err != nil {
			fmt.Println(err)
			w.Write([]byte("wrong credentials"))
			return
		}

		w.Write([]byte("ok. your id is " + strconv.Itoa(id)))
	})

	http.ListenAndServe(":80", nil)
}

// http://localhost/login?username=sulaev&password=123123
// http://localhost/login?username=admin&password="OR"1"="1
