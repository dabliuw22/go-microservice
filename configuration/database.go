package configuration

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	database = "my_db"
)

const driverName  = "postgres"

func dataSource() string {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)
	return dataSource
}

func connection(dataSource string) (*gorm.DB, error) {
	return gorm.Open(driverName, dataSource)
}