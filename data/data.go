package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./configs.db")

	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateSupportedIntegrationsTable() {
	createSupportedIntegrationsTableSQL := `CREATE TABLE IF NOT EXISTS supported_integrations (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "integration_type" TEXT
  );`

	statement, err := db.Prepare(createSupportedIntegrationsTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	//log.Println("Integrations are initialized")
}

func CreateIntegrationsTable() {
	createIntegrationsTableSQL := `CREATE TABLE IF NOT EXISTS integrations (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "integration_type" TEXT,
	"name" TEXT,
    "apikey" TEXT
  );`

	statement, err := db.Prepare(createIntegrationsTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("Integrations are initialized")
}

func InsertIntegration(integration_type string, name string, apiKey string) {
	insertIntegrationSQL := `INSERT INTO integrations(integration_type, name, apikey) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertIntegrationSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(integration_type, name, apiKey)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Integration saved successfully")

}

func InsertSupportedIntegration(integration_type string) {
	insertIntegrationSQL := `INSERT INTO supported_integrations(integration_type) VALUES (?)`
	statement, err := db.Prepare(insertIntegrationSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(integration_type)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Integration saved successfully")

}

func RemoveIntegration(id int64) {
	deleteIntegrationSQL := `DELETE FROM integrations WHERE id = ?`
	statement, err := db.Prepare(deleteIntegrationSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(id)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Integration removed successfully")

}

func GetIntegrations() []Integration {
	rows, err := db.Query("SELECT * FROM integrations")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var integrations []Integration
	for rows.Next() {
		var id int
		var integration_type string
		var name string
		var apiKey string
		err = rows.Scan(&id, &integration_type, &name, &apiKey)
		if err != nil {
			log.Fatal(err)
		}
		integrations = append(integrations, Integration{id, integration_type, name, apiKey})
	}
	return integrations
}

func GetSupportedIntegrations() []SupportedIntegration {
	rows, err := db.Query("SELECT * FROM supported_integrations")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var supported_integrations []SupportedIntegration
	for rows.Next() {
		var id int
		var integration_type string
		err = rows.Scan(&id, &integration_type)
		if err != nil {
			log.Fatal(err)
		}
		supported_integrations = append(supported_integrations, SupportedIntegration{id, integration_type})
	}
	return supported_integrations
}

func PrintIntegrations(integrations []Integration) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Integration Type", "Name", "API Key"})

	for _, i := range integrations {
		table.Append([]string{fmt.Sprintf("%d", i.Id), i.IntegrationType, i.Name, MaskAPIKey(i.ApiKey)})
	}

	table.SetBorder(false)        // Disable default border
	table.SetCenterSeparator("|") // Separator for better readability
	table.SetColumnSeparator("|") // Vertical separator
	table.SetRowSeparator("-")    // Horizontal separator
	table.Render()                // Render the table
}

func PrintSupportedIntegrations(supported_integrations []SupportedIntegration) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Supported Integration"})

	for _, i := range supported_integrations {
		table.Append([]string{fmt.Sprintf("%d", i.Id), i.SupIntegration})
	}

	table.SetBorder(false)        // Disable default border
	table.SetCenterSeparator("|") // Separator for better readability
	table.SetColumnSeparator("|") // Vertical separator
	table.SetRowSeparator("-")    // Horizontal separator
	table.Render()                // Render the table
}

func MaskAPIKey(apiKey string) string {
	if len(apiKey) > 4 {
		return "****" + apiKey[len(apiKey)-4:]
	}
	return "****"
}

type SupportedIntegration struct {
	Id             int
	SupIntegration string
}

type Integration struct {
	Id              int
	IntegrationType string
	Name            string
	ApiKey          string
}
