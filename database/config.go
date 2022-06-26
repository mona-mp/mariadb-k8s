package database

import (
	"fmt"
)

//database server config
type Config struct {
	ServerName string
	User       string
	Password   string
	DBName     string
}

var GetConnectionString = func(config Config) string {
	//create connection string to connect to mysql
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.User, config.Password, config.ServerName, config.DBName)

	return connectionString
}
