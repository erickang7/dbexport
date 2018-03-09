package exportcsv

import "fmt"

// GenerateConnectionString function generates mssql connection string based on *Config struct
func GenerateConnectionString(c *Config) string {
	return fmt.Sprintf("server=%s; initial catalog=%s; user id=%s;password=%s;port=%d",
		c.ServerName, c.DatabaseName, c.User, c.Password, c.Port)
}
