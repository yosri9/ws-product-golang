package database

import (
	"database/sql"
	"log"
	"server/server/env"
)

var (
	dataSourceName = env.User + ":" + env.Password + "@tcp(" + env.DatabaseAddress + ":" + env.DatabasePort + ")/" + env.DatabaseName
)

func DatabaseConnection() sql.DB {
	// Capture connection properties.
	db, err := sql.Open(env.DriverName, dataSourceName)

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}

	defer db.Close()

	return *db
}
