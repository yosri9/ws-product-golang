package main

import (
	"server/server/database/table"
	"server/server/model"
	"server/server/route"
)

func main() {
	table.CreateCounterTable()
	model.MapCounter = make(map[string]string)
	route.RouteHandler()

}
