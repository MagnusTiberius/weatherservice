package handler

import (
	"net/http"
)

type ResponseMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	name := q.Get("name")
	if name == "" {
		name = "unknown"
	}

	response := ResponseMessage{Message: "Hello, " + name, Code: 200}
	JsonResponseWrite(w, response, 200)
}
