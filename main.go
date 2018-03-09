package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/erickang7/dbexport/exportcsv"
)

const configPath = "./dbconfig.yaml"

var config exportcsv.Config

//var db *sql.DB

func main() {

	if config.LoadConfig(configPath) != nil {
		log.Fatal("failed to load connection configuration file")
	}

	connString := exportcsv.GenerateConnectionString(&config)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected to\n")
	log.Print(connString)

	defer db.Close()
	ctx := context.Background()
	var result string
	err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("%s\n", result)

}
