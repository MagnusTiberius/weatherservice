package main

import (
	"net/http"

	"github.com/MagnusTiberius/weatherservice/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.IndexHandler)
	r.HandleFunc("/user/{username}", handler.UserHandler)
	http.ListenAndServe(":8080", r)
}
