package route

import (
	"log"
	"net/http"
	"server/server/env"
	"server/server/handler"
	"server/server/middleware/per_client_middleware"
)

var (
	mux = http.NewServeMux()

	w http.ResponseWriter
)

func RouteHandler() {

	statHandler, middleware := per_client_middleware.StatMiddleware()
	mux.HandleFunc("/", handler.WelcomeHandler)
	mux.HandleFunc("/view/", handler.ViewHandler)
	mux.Handle("/stats/", middleware.Handle(statHandler))

	err := http.ListenAndServe(env.IPAddress+":"+env.Port, mux)
	log.Fatal(err)

}
