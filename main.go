package main

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func init() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=ex password=example dbname=rain sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("postgre db init success")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("postgre db ping success")
}

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Signin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
