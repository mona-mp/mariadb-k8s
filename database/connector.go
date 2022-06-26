package database

import (
	"log"
	"user-management/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

var Connector *gorm.DB

func Connect(connectionString string) error {

	var err error
	Connector, err = gorm.Open("mysql", connectionString)

	if err != nil {
		return err
	}
	log.Println("Connection was successful!")
	return nil
}

func Migrate(table *entity.User) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
