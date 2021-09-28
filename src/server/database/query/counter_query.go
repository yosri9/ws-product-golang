package query

import (
	"server/server/database"
	"server/server/model"
)

var (
	db = database.DatabaseConnection()
)

func Create(c *model.Counters) {
	_, err := db.Query("INSERT INTO `product`.`counters` (`event_name`, `event_time`, `view_count`, `click_count`) VALUES (?, ?, ?, ?);", c.Name, c.Time, c.View, c.Click)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app

	}

	return
}

func Update(c *model.Counters) {
	_, err := db.Query("UPDATE counters SET view_count = ?, click_count = ? WHERE (event_name = ? and event_time = ?);", c.View, c.Click, c.Name, c.Time)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func Exist(c *model.Counters) bool {
	dbCounters, err := db.Query("Select * from counters where (event_name = ? and event_time=?) ", c.Name, c.Time)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	print(dbCounters.Next())

	return dbCounters.Next()
}

func delete(id int) {

}

func getOne(id int) {

}
func getAll() {

}
