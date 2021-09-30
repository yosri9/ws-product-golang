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
		view(data, key, value, w)

	case "entertainment":
		model.EntertainmentCounter.Time = time
		model.CurrentCounter = &(model.EntertainmentCounter)
		view(data, key, value, w)

	case "business":
		model.BusinessCounter.Time = time
		model.CurrentCounter = &(model.BusinessCounter)
		view(data, key, value, w)

	case "education":
		model.EducationCounter.Time = time
		model.CurrentCounter = &(model.EducationCounter)
		view(data, key, value, w)

	}

	// simulate random click call
	if rand.Intn(100) < 50 {
		process.ProcessClick(model.CurrentCounter)
	}

	err := process.ProcessRequest(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

}

func view(data string, key string, value string, w http.ResponseWriter) {
	go process.ProcessView(model.CurrentCounter, data, key)
	go controller.AppendToCountersList(model.CurrentCounter)
	fmt.Fprint(w, key, value)
}
