package table

import "server/server/database"

func CreateCounterTable() {
	db := database.DatabaseConnection()
	//db.Ping()
	//db.Exec("CREATE TABLE IF NOT EXISTS `counters` `id` BIGINT NOT NULL auto_increment, `event_name` varchar(250)  NOT NULL default, `event_time` varchar(250)  NOT NULL default, `event_click` int, `event_view` int",
	//	)
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `counters` " +
		"(\n\t`id` INT NOT NULL AUTO_INCREMENT,\n\t`event_name` VARCHAR(50) NULL DEFAULT NULL," +
		"\n\t`event_time` DATETIME NULL DEFAULT NULL,\n\t`view_count` INT NULL DEFAULT NULL," +
		"\n\t`click_count` INT NULL DEFAULT NULL,\n\tPRIMARY KEY (`id`)\n);")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//ToDo counterQuery
	//ToDo convert date format(y-m-d h-m)
}
