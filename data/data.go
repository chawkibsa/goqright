package data

import (
 "database/sql"
 _ "github.com/mattn/go-sqlite3"
 "log"
)

var db *sql.DB

func OpenDatabase() error {
  var err error

  db, err = sql.Open("sqlite3","./configs.db")

  if err != nil {
    return err
  }

  return db.Ping()
}

func CreateIntegrationsTable() {
  createIntegrationsTableSQL := `CREATE TABLE IF NOT EXISTS integrations (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "integration" TEXT,
    "api" TEXT
  );`

  statement, err := db.Prepare(createIntegrationsTableSQL)
  if err != nil {
    log.Fatal(err.Error())
  }

  statement.Exec()
  log.Println("Integrations are initialized")
}
