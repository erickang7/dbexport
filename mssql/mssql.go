package mssql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/sqltocsv"
)

// GenerateConnectionString generates mssql connection string based on *Config struct
func generateConnectionString(c *Config) string {
	return fmt.Sprintf("server=%s; database=%s; user id=%s;password=%s;port=%d",
		c.ServerName, c.DatabaseName, c.User, c.Password, c.Port)
}

//Connect to a sql server instance
func Connect(c *Config) (*sql.DB, error) {

	db, err := sql.Open("sqlserver", generateConnectionString(c))
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	return db, nil
}

// ExecuteQuery ..
func ExecuteQuery(db *sql.DB, query string) (*sql.Rows, error) {
	ctx := context.Background()
	return db.QueryContext(ctx, query)
}

// GetServerVersion ...
func GetServerVersion(db *sql.DB) (*sql.Rows, error) {
	var query = "SELECT @@version"
	return ExecuteQuery(db, query)
}

//GetTableList ...
func GetTableList(db *sql.DB) ([]string, error) {
	var tableList []string
	var query = "select concat(table_schema, '.', table_name) as tables "
	query += "from information_schema.tables where table_type = N'BASE TABLE'"

	rows, err := ExecuteQuery(db, query)
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		tableList = append(tableList, tableName)
	}
	return tableList, err
}

//PrintRows ...
func PrintRows(rows *sql.Rows) {
	for rows.Next() {
		var row string
		rows.Scan(&row)
		fmt.Println(row)
	}
}

// SaveAsCSV ...
func SaveAsCSV(c *Config, db *sql.DB, filePath []string) {

	for _, file := range filePath {
		var query = "SELECT * FROM " + file
		rows, err := ExecuteQuery(db, query)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("Exporting %s to CSV...\n", file)
		fileName := buildCSVPath(c.CSVPath, file)
		isFile(fileName)
		if sqltocsv.WriteFile(fileName, rows) != nil {
			log.Fatal(err.Error())
		}
		rows.Close()
	}
}

func buildCSVPath(dirPath string, fileName string) string {
	return dirPath + fileName + ".csv"
}

func isFile(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		fmt.Printf("%s created\n", filename)
	}
}
