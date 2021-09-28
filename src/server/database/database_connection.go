package database

import (
	"database/sql"
	"log"
)

type person struct {
	firstName string
	lastName  string
}

func DatabaseConnection() sql.DB {
	// Capture connection properties.
	db, err := sql.Open(driverName, dataSourceName)

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}

	defer db.Close()

	//Execute the query
	//results, err := db.Query("SELECT first_name, last_name FROM users")
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//for results.Next() {
	//	var  person person
	//	err = results.Scan(&person.firstName, &person.lastName)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	log.Printf(person.firstName)
	//	print("hello brother\n")
	//	print(person.firstName + "\n")
	//	print(person.lastName + "\n")
	//
	//}

	//for results.Next() {
	//
	//	var tag Tag
	//	// for each row, scan the result into our tag composite object
	//	err = results.Scan(&tag.ID, &tag.Name)
	//	if err != nil {
	//		panic(err.Error()) // proper error handling instead of panic in your app
	//	}
	//	//and then print out the tag's Name attribute
	//	log.Printf(tag.Name)
	//}

	return *db
}
