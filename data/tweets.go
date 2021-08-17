package data

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"time"
)

// basic tweet information
type Tweet struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	CreatedOn string `json:"created"`
}

// Tweets is a collection of Tweet
type Tweets []Tweet

// tweets to JSON
func (t *Tweets) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}

// JSON to tweet
func (t *Tweet) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}

// adds a new tweet to our collection of tweets
func CreateTweet(t *Tweet, db *sql.DB) {

	// generate the timestamp
	t.CreatedOn = time.Now().UTC().String()

	// insert tweet into db, ID auto created
	querySql := `
		INSERT INTO tweets (text, createdon)
		VALUES ($1, $2)`

	_, err := db.Exec(querySql, t.Text, t.CreatedOn)
	if err != nil {
		log.Fatal(err)
	}
}

// GetTweets return tweets
func GetTweets(db *sql.DB, l *log.Logger) Tweets {
	// get tweets from db here
	rows, err := db.Query("SELECT id, text, createdon from tweets;")
	if err != nil {
		l.Fatal("Error pulling from tweets table", err)
	}
	defer rows.Close()

	var tw Tweets

	for rows.Next() {
		var t Tweet

		err := rows.Scan(&t.ID, &t.Text, &t.CreatedOn)
		if err != nil {
			l.Fatal("Error while scanning rows", err)
		}
		tw = append(tw, t)
	}

	return tw
}

func CreateTable(db *sql.DB, l *log.Logger) {
	query, err := db.Prepare(`CREATE TABLE IF NOT EXISTS tweets (
		id INTEGER PRIMARY KEY, 
		text TEXT, 
		createdon TEXT
		)`,
	)
	if err != nil {
		l.Fatal("Error creating table: ", err)
	}
	query.Exec()
}
