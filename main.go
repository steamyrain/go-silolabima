package main

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	r.GET("/signin", Signin)
	r.GET("/welcome", Welcome)
	r.Run()
}

func Signin(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Hello, %q", html.EscapeString(c.Request.URL.Path))
}

func Welcome(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Hello, %q", html.EscapeString(c.Request.URL.Path))
}
