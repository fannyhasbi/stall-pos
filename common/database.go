package common

import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/jmoiron/sqlx"
  "log"
  "fmt"
)

const (
  user = "root"
  pass = ""
  dbname = "stall_pos"
)

func Connect() (*sqlx.DB, error) {
  var conn string
  if len(pass) > 0 {
    conn = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", user, pass, dbname)
  } else {
    conn = fmt.Sprintf("%s@tcp(localhost:3306)/%s", user, dbname)
  }

  db, err := sqlx.Open("mysql", conn)
	if err != nil {
    log.Fatal(err)
    return db, err
	}

	return db, nil
}