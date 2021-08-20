package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"example.com/main/data"
	"example.com/main/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	var db *sql.DB
	var err error

	l := log.New(os.Stdout, "tweets-api ", log.LstdFlags)

	// open db and defer close until execute
	db, err = sql.Open("sqlite3", os.Getenv("sqldatabase"))
	if err != nil {
		l.Fatal("error establishing db connection", err)
	}
	defer db.Close()

	// creates table if it doesn't exist
	data.CreateTable(db, l)

	// create the handlers
	th := handlers.NewTweets(l, db)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter().StrictSlash(true)
	sm.Handle("/tweets", th)

	// log and start server, fatal shut downs server
	l.Println("Starting server on port 9090")
	l.Fatal(http.ListenAndServe(":9090", sm))

}
