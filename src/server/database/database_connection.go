package database

import (
	"database/sql"
	"log"
)

var (
	dataSourceName = user + ":" + password + "@tcp(" + address + ":" + port + ")/" + database
)

func DatabaseConnection() sql.DB {
	// Capture connection properties.
	db, err := sql.Open(driverName, dataSourceName)

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}

	defer db.Close()

	return *db
}
