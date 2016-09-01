package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysql = "user:password@tcp(127.0.0.1:3306)/fofgaming"

type mysqlBackend struct {
	db *sql.DB
}

func (m *mysqlBackend) sawUser(userID string, now time.Time) {
	s, err := m.db.Prepare("INSERT INTO `members` (`slack`,`seen`) VALUES(?,?) ON DUPLICATE KEY UPDATE `seen`=VALUES(`seen`)")
	if err != nil {
		log.Fatal("Error preparing query for MySQL", err)
	}
	_, err = s.Exec(userID, now.Unix())
	if err != nil {
		log.Fatal("Error updating MySQL", err)
	}
}

func (m *mysqlBackend) connect() {
	db, err := sql.Open("mysql", mysql)
	if err != nil {
		log.Fatal("Error connecting to MySQL", err)
	}
	m.db = db
}

func (m *mysqlBackend) close() {
	m.db.Close()
}

func init() {
	if connInfo := os.Getenv("SEEN_MYSQL_URI"); connInfo != "" {
		mysql = connInfo
	}
	flag.StringVar(&mysql, "mysql", mysql, "MySQL Connection URI [ SEEN_MYSQL_URI ]")
}

func initMySQL() {
	var backend = new(mysqlBackend)
	backend.connect()
	seen.register(backend)
}
