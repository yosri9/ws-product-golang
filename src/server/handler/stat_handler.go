package handler

import (
	"encoding/json"
	"net/http"
	"server/server/database/query"
	"server/server/middleware/global_middleware"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	go global_middleware.GlobalMiddleware()
	w.Header().Set("Content-Type", "application/json")

	// post all counter in json form
	for _, counter := range query.GetAll() {
		json.NewEncoder(w).Encode(counter)

	}

	json.NewEncoder(w).Encode(query.GetOne("sport", "2021-09-28 09:06"))

}
