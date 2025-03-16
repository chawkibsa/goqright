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
    "apikey" TEXT
  );`

  statement, err := db.Prepare(createIntegrationsTableSQL)
  if err != nil {
    log.Fatal(err.Error())
  }

  statement.Exec()
  log.Println("Integrations are initialized")
}

func InsertIntegration(integration string, apiKey string) {
  insertIntegrationSQL := `INSERT INTO integrations(integration, apikey) VALUES (?, ?)`
  statement, err := db.Prepare(insertIntegrationSQL)
  if err != nil {
    log.Fatalln(err.Error())
  }

  _, err = statement.Exec(integration, apiKey)
  if err != nil {
    log.Fatalln(err.Error())
  }

  log.Println("Integration inserted successfully")

}
