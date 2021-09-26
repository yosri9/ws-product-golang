package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"server/model"
	"server/process"
	"time"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	//get random value from content list
	data := model.Content[rand.Intn(len(model.Content))]
	key := data + ": " + string(time.Now().Format("2006.01.02 15:04"))
	value := model.MapCounter[key]

	switch data {
	case "sports":
		model.CurrentCounter = &(model.SportCounter)
		process.ProcessView(&(model.SportCounter), data, key)
		fmt.Fprint(w, string(key), string(value))

	case "entertainment":
		model.CurrentCounter = &(model.EntertainmentCounter)
		process.ProcessView(&model.EntertainmentCounter, data, key)
		fmt.Fprint(w, string(key), string(value))

	case "business":
		model.CurrentCounter = &(model.BusinessCounter)
		process.ProcessView(&(model.BusinessCounter), data, key)
		fmt.Fprint(w, string(key), string(value))

	case "education":
		model.CurrentCounter = &(model.EducationCounter)
		process.ProcessView(&(model.EducationCounter), data, key)
		fmt.Fprint(w, string(key), string(value))
	}
	// simulate random click call
	if rand.Intn(100) < 50 {
		process.ProcessClick(model.CurrentCounter, data)
	}

	//c.Lock()
	//c.view++
	//c.Unlock()

	err := process.ProcessRequest(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

}
