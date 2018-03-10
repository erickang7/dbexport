package main

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/erickang7/dbexport/mssql"
)

const configPath = "./dbconfig.yaml"

var config mssql.Config
var db *sql.DB

func main() {

	// load database connetion configuration yaml
	if config.LoadConfig(configPath) != nil {
		log.Fatal("failed to load connection configuration file")
	}

	// open a connection
	db, err := mssql.Connect(&config)
	if err != nil {
		log.Fatalf("connection failed to %s: ", config.ServerName)
	}

	// optional: validate connection
	result, err := mssql.GetServerVersion(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	mssql.PrintRows(result)
	result.Close()

	// get table names
	tableList, err := mssql.GetTableList(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	// export tables to CSV
	mssql.SaveAsCSV(&config, db, tableList)

	defer db.Close()

}
