package main

import (
	"database/sql"
	"flag"
	"log"
	"time"
)

var dbPath = "./db.sqlite3"

type sqliteBackend struct {
	db          *sql.DB
	updateQuery *sql.Stmt
}

func (s *sqliteBackend) sawUser(userID string, now time.Time) {
	if _, err := s.updateQuery.Exec(userID, now.String()); err != nil {
		log.Println(err)
	}
}

func (s *sqliteBackend) open() {
	if db, err := sql.Open("sqlite3", dbPath); err != nil {
		log.Fatal(err)
	} else {
		s.db = db
	}
}

func (s *sqliteBackend) prepare() {
	if _, err := s.db.Exec("CREATE TABLE IF NOT EXISTS `seen` (`id` STRING PRIMARY KEY,`when` TEXT);"); err != nil {
		log.Fatal(err)
	}

	if stmt, err := s.db.Prepare("INSERT OR REPLACE INTO `seen` (`id`,`when`) VALUES(?,?)"); err != nil {
		log.Fatal(err)
	} else {
		s.updateQuery = stmt
	}
}

func (s *sqliteBackend) load() {
	if rows, err := s.db.Query("SELECT * FROM `seen`"); err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			var id string
			var when string
			rows.Scan(&id, &when)
			if id == "" {
				continue
			}
			seen[id], err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", when)
			if err != nil {
				log.Println("failed parsing time", when, "for", id)
			}
		}
		rows.Close()
	}
	for u, t := range seen {
		log.Println("Loaded", u, "last seen time", t)
	}
}

func init() {
	flag.StringVar(&dbPath, "db", dbPath, "path to the database")
}

func initSqlite() {
	var backend = new(sqliteBackend)
	backend.open()
	backend.prepare()
	backend.load()
	seen.register(backend)
}
