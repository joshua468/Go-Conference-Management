package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "./conference.db")
	if err != nil {
		log.Fatal(err)
	}
	createTables()
	return db
}

func createTables() {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS attendees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);

	CREATE TABLE IF NOT EXISTS conferences (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);

	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

type Attendee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateAttendee(attendee *Attendee) error {
	_, err := db.Exec("INSERT INTO attendees (name) VALUES (?)", attendee.Name)
	return err
}

type Conference struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateConference(conference *Conference) error {
	_, err := db.Exec("INSERT INTO conferences (name) VALUES (?)", conference.Name)
	return err
}

type Session struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateSession(session *Session) error {
	_, err := db.Exec("INSERT INTO sessions (title, description) VALUES (?, ?)", session.Title, session.Description)
	return err
}
