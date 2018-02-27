package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	ApiAddr = "35.225.142.51:8090"
)

type ResponseMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/address/{addr}", AddressHandler)
	r.HandleFunc("/address/{addr}/{time}", AddressHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8088", "http://35.226.247.163:8088/"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	http.ListenAndServe(":8090", handler)
}

func JsonResponseWrite(w http.ResponseWriter, message interface{}, statusCode int) {

	body, err := json.Marshal(message)

	if statusCode == 200 && err == nil {
		msg, _ := json.Marshal(message)
		w.Header().Set("content-type", "application/json")
		w.Write(msg)
	} else {
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			http.Error(w, string(body), statusCode)
		}
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	name := q.Get("name")
	if name == "" {
		name = "unknown"
	}

	response := ResponseMessage{Message: "", Code: 200}
	JsonResponseWrite(w, response, 200)
}

func AddressHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	addr := vars["addr"]
	if addr == "" {
		addr = "unknown"
		log.Print("addr undefined")
		response := ResponseMessage{Message: "", Code: 200}
		JsonResponseWrite(w, response, 200)
	}
	tm := vars["time"]
	if tm == "" {
		log.Print("time undefined")
	}

	urlApi := ""
	if tm == "" {
		urlApi = fmt.Sprintf("http://%s/address/%s", ApiAddr, addr)
	} else {
		urlApi = fmt.Sprintf("http://%s/address/%s/%s", ApiAddr, addr, tm)
	}
	log.Printf("urlApi:%s\n\n", urlApi)

	response, err := http.Get(urlApi)
	if err != nil {
		log.Print(err.Error())
		response := ResponseMessage{Message: "", Code: 200}
		JsonResponseWrite(w, response, 200)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		response := ResponseMessage{Message: "", Code: 200}
		JsonResponseWrite(w, response, 200)
	}

	w.Header().Set("content-type", "application/json")
	w.Write(responseData)
}
