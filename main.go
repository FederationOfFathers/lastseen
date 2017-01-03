package main

import (
	"flag"
)

func main() {
	defer seen.close()
	flag.Parse()
	initSqlite()
	initMySQL()
	mindHTTP()
	mindSlack()
}
