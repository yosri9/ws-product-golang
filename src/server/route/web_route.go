package route

import (
	"log"
	"net/http"
	"server/server/handler"
	"server/server/middleware"
)

var (
	mux = http.NewServeMux()
)

func RouteHandler() {
	statHandler, middleware := middleware.StatMiddleware()
	mux.HandleFunc("/", handler.WelcomeHandler)
	mux.HandleFunc("/view/", handler.ViewHandler)
	mux.Handle("/stats/", middleware.Handle(statHandler))

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
