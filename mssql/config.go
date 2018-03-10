package mssql

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Config stuct maps to dbconfig.yaml
type Config struct {
	ServerName   string `jason:"servername" yaml:"servername"`
	Port         int    `jason:"port" yaml:"port"`
	DatabaseName string `jason:"databasename" yaml:"databasename"`
	User         string `jason:"user" yaml:"user"`
	Password     string `jason:"password" yaml:"password"`
	CSVPath      string `jason:"csvpath" yaml:"csvpath"`
}

// LoadConfig method reads dbconfig.yaml file and instantiates *Config struct
func (c *Config) LoadConfig(configPath string) error {
	var err error
	var y []byte
	y, err = ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("failed to read ", configPath)
		return err
	}

	if err = yaml.Unmarshal(y, &c); err != nil {
		log.Fatal("failed to unmarshar Config")
		return err
	}
	return nil
}

// String() function prints out Config struct property values
func (c *Config) String() string {
	return fmt.Sprintf("server: %s\nport: %d\ndatabase: %s\nuser: %s\n",
		c.ServerName, c.Port, c.DatabaseName, c.User)
}
