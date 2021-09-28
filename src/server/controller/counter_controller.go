package controller

import (
	"server/server/database/query"
	"server/server/model"
)

var (
	counters []model.Counter
)

func countTracking(event string, time string) string {
	key := event + ":" + time
	return model.MapCounter[key]
}

func CountCreateOrUpdate(c *model.Counter) {
	if query.Exist(c) {
		query.Update(c)
	} else {
		query.Create(c)
	}
	return
}

func countDelete() {

}

func find() {

}
func findOne() {

}

func isAllowed() bool {

	return true
}

func UploadCounters() error {
	for _, counter := range counters {
		CountCreateOrUpdate(&counter)
	}
	// free counter list from value every 5 seconds
	counters = counters[:0]

	return nil
}

func AppendToCountersList(c *model.Counter) {
	counters = append(counters, *c)

}
