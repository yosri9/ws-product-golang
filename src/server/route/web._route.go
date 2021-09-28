package route

import (
	"log"
	"net/http"
	"server/server/handler"
)

func RouteHandler() {
	http.HandleFunc("/", handler.WelcomeHandler)
	http.HandleFunc("/view/", handler.ViewHandler)
	http.HandleFunc("/stats/", handler.StatsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
