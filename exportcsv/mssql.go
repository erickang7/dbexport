package exportcsv

import "fmt"

// GenerateConnectionString ...
func GenerateConnectionString(c *Config) string {
	return fmt.Sprintf("server=%s; initial catalog=%s; user id=%s;password=%s;port=%d",
		c.ServerName, c.DatabaseName, c.User, c.Password, c.Port)
}
