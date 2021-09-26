package main

import (
	"log"
	"net/http"
	"server/handler"
	"server/model"
)

//func isAllowed() bool {
//
//	return true
//}
//
//func uploadCounters() error {
//	return nil
//}
//

func main() {
	model.MapCounter = make(map[string]string)

	http.HandleFunc("/", handler.WelcomeHandler)
	http.HandleFunc("/view/", handler.ViewHandler)
	http.HandleFunc("/stats/", handler.StatsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
