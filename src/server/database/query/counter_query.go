package query

import (
	"server/server/database"
	"server/server/model"
)

var (
	db = database.DatabaseConnection()
)

func Create(c *model.Counter) {
	//Execute the query
	_, err := db.Query("INSERT INTO `product`.`counters` (`event_name`, `event_time`, `view_count`, `click_count`) VALUES (?, ?, ?, ?);", c.Name, c.Time, c.View, c.Click)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app

	}

	return
}

func Update(c *model.Counter) {
	_, err := db.Query("UPDATE counters SET view_count = ?, click_count = ? WHERE (event_name = ? and event_time = ?);", c.View, c.Click, c.Name, c.Time)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func Exist(c *model.Counter) bool {
	dbCounters, err := db.Query("Select * from counters where (event_name = ? and event_time=?) ", c.Name, c.Time)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	print(dbCounters.Next())

	return dbCounters.Next()
}

func delete(id int) {

}

func GetOne(eventName string, eventTime string) model.Counter {
	var counter model.Counter

	dbCounters, err := db.Query("Select * from counters where  (event_name=? and event_time=?)", eventName, eventTime)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if dbCounters.Next() {

		err = dbCounters.Scan(&counter.ID, &counter.Name, &counter.Time, &counter.View, &counter.Click)
		if err != nil {
			panic(err.Error())
		}

	}
	return counter

}
func GetAll() []model.Counter {
	counters := []model.Counter{}
	var counter model.Counter

	//Execute the query
	dbCounters, err := db.Query("Select * from counters ")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for dbCounters.Next() {
		err = dbCounters.Scan(&counter.ID, &counter.Name, &counter.Time, &counter.View, &counter.Click)
		//ToDo remove sec from time
		if err != nil {
			panic(err.Error())
		}
		counters = append(counters, counter)
		print(len(counters))

	}

	return counters

}
