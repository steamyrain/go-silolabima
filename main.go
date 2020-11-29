package main

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func init() {
	db, err := sql.Open("postgres", os.Getenv("ALKAL_DB_URI"))
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
