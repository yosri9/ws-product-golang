package handler

import (
	"encoding/json"
	"net/http"
	"server/server/database/query"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// post all counter in json form
	for _, counter := range query.GetAll() {
		json.NewEncoder(w).Encode(counter)

	}

	json.NewEncoder(w).Encode(query.GetOne("sport", "2021-09-28 09:06"))

}
