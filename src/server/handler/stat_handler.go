package handler

import (
	"encoding/json"
	"net/http"
	"server/server/database/query"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//
	//for _, counter := range query.GetAll() {
	//	json.NewEncoder(w).Encode(counter)
	//
	//	//print(counter.Name)
	//	//fmt.Fprint(w, counter.Name, ":" ,counter.Time,"  " , "{view: ", counter.View, " click: ", counter.Click  )
	//
	//}

	json.NewEncoder(w).Encode(query.GetOne("sport", "2021-09-28 09:06"))

	//if !isAllowed() {
	//	w.WriteHeader(429)
	//	return
	//}
}
