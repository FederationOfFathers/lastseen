package main

import (
	"flag"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	flag.Parse()
	initSqlite()
	mindSlack()
	mindHTTP()
}
