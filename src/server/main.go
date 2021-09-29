package main

import (
	"server/server/controller"
	"server/server/database/query"
	"server/server/database/table"
	"server/server/model"
	"server/server/route"
)

func main() {
	model.MapCounter = make(map[string]string)

	go table.CreateCounterTable()
	go query.GetAll()
	go controller.UploadCountersEvery5Sec()
	route.RouteHandler()

	//route.RouteHandler()

}
