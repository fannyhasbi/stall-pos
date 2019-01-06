package common

import (
  "database/sql"
  "log"
  "fmt"
)

const (
  user = "root"
  pass = ""
  dbname = "stall_pos"
)

func Connect() *sql.DB {
  var conn string
  if len(conn) > 0 {
    conn = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", user, pass, dbname)
  } else {
    conn = fmt.Sprintf("%s@tcp(localhost:3306)/%s", user, dbname)
  }

  db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Fatal(err)
	}

	return db
}