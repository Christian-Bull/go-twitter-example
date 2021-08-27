package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"example.com/main/data"
)

// Tweets handler
type Tweets struct {
	l  *log.Logger
	db *sql.DB
}

// New tweet handler
func NewTweets(l *log.Logger, db *sql.DB) *Tweets {
	return &Tweets{l, db}
}

// Handler interface
func (t *Tweets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle request for method
	if r.Method == http.MethodGet {
		// sets return type as json
		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		t.getTweets(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		t.createTweet(rw, r)

		return
	}
}

// get tweets
func (t *Tweets) getTweets(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle get tweets")

	lp := data.GetTweets(t.db, t.l)

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error marshalling json", http.StatusInternalServerError)
	}
}

// create tweet
func (t *Tweets) createTweet(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Handle POST Tweet")

	d := &data.Tweet{}

	err := d.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Error unmarshalling json", http.StatusBadRequest)
	}

	err = data.CreateTweet(d, t.db, t.l)
	if err != nil {
		http.Error(rw, "Error creating tweet", http.StatusInternalServerError)
	}

	// set response header
	rw.WriteHeader(http.StatusCreated)
}
