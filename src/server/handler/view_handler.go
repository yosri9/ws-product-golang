package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"server/server/controller"
	"server/server/model"
	"server/server/process"
	"time"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	//get random value from content list
	data := model.Content[rand.Intn(len(model.Content))]
	// key will be like "sports:2020-01-08 22:01"
	time := time.Now().Format("2006.01.02 15:04")
	key := data + ": " + time
	value := model.MapCounter[key]

	//update or call to  counter depending on the event name

	switch data {
	case "sports":
		model.SportCounter.Time = time
		model.CurrentCounter = &(model.SportCounter)
		process.ProcessView(&(model.SportCounter), data, key)
		controller.AppendToCountersList(model.CurrentCounter)
		fmt.Fprint(w, string(key), string(value))

	case "entertainment":
		model.EntertainmentCounter.Time = time
		model.CurrentCounter = &(model.EntertainmentCounter)
		process.ProcessView(&model.EntertainmentCounter, data, key)
		controller.AppendToCountersList(model.CurrentCounter)
		fmt.Fprint(w, string(key), string(value))

	case "business":
		model.BusinessCounter.Time = time
		model.CurrentCounter = &(model.BusinessCounter)
		process.ProcessView(&(model.BusinessCounter), data, key)
		controller.AppendToCountersList(model.CurrentCounter)
		fmt.Fprint(w, string(key), string(value))

	case "education":
		model.EducationCounter.Time = time
		model.CurrentCounter = &(model.EducationCounter)
		controller.AppendToCountersList(model.CurrentCounter)
		process.ProcessView(&(model.EducationCounter), data, key)
		fmt.Fprint(w, string(key), string(value))
	}

	// simulate random click call
	if rand.Intn(100) < 50 {
		process.ProcessClick(model.CurrentCounter, data)
	}

	err := process.ProcessRequest(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

}
