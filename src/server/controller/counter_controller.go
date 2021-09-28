package controller

import (
	"math/rand"
	"server/server/database/query"
	"server/server/model"
	"time"
)

func countTracking(event string, time string) string {
	key := event + ":" + time
	return model.MapCounter[key]
}

func CountCreateOrUpdate(c *model.Counters) {
	time.Sleep(time.Duration(rand.Int31n(50)) * time.Millisecond)
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

func uploadCounters() error {
	return nil
}
