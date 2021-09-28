package main

import (
	"server/server/database/query"
	"server/server/database/table"
	"server/server/model"
	"server/server/route"
)

func main() {
	query.GetAll()
	table.CreateCounterTable()
	model.MapCounter = make(map[string]string)
	route.RouteHandler()

}
