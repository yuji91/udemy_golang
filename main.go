package main

import (
  "database/sql"
  "log"

  _ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func main() {
  Db, _ := sql.Open("sqlite3", "./example.sql")

  defer Db.Close()

  // INSERT DATA
  cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
  _, err := Db.Exec(cmd, "test_01", 20)
  if err != nil {
    log.Fatalln(err)
  }

/*
  // CREATE TABLE
  cmd := `CREATE TABLE IF NOT EXISTS persons(
                                      name STRING,
                                      age INT)`
  _, err := Db.Exec(cmd)

  if err != nil {
    log.Fatalln(err)
  }
*/
}
