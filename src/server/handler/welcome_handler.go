package handler

import (
	"fmt"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to EQ Works 😎")
}
