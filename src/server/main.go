package main

import (
	"server/server/controller"
	"server/server/database/query"
	"server/server/database/table"
	"server/server/model"
	"server/server/route"
	"time"
)

func main() {
	model.MapCounter = make(map[string]string)

	go table.CreateCounterTable()
	go query.GetAll()
	go uploadCountersEvery5Sec()

	route.RouteHandler()

}

func uploadCountersEvery5Sec() {

	uptimeTicker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-uptimeTicker.C:
			controller.UploadCounters()
		}
	}
}
