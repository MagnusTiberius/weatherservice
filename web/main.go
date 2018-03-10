package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var tmpl *template.Template

type ResponseMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Address struct {
	addr string
}

func main() {
	tmpl = template.Must(template.ParseFiles("/var/www/html/index.html"))
	//tmpl = template.Must(template.ParseFiles("web/index.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/address/{addr}", AddressHandler)
	r.HandleFunc("/address/{addr}/{time}", AddressHandler)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8089", "http://35.226.247.163:8089/"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	http.ListenAndServe(":8089", handler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	name := q.Get("name")
	if name == "" {
		name = "unknown"
	}
	tmpl.Execute(w, struct{ Success bool }{true})
}

func AddressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := ""
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
		url = "http://35.226.125.10:8091/address/" + addr
	} else {
		url = "http://35.226.125.10:8091/address/" + addr + "/" + tm
	}

	fmt.Printf("url:%s\n\n", url)
	response, err := http.Get(url)
	if err != nil {
		log.Print(err.Error())
		resp := ResponseMessage{Message: "", Code: 200}
		JsonResponseWrite(w, resp, 200)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err.Error())
		response := ResponseMessage{Message: "", Code: 200}
		JsonResponseWrite(w, response, 200)
	}
	w.Header().Set("content-type", "application/json")
	w.Write(responseData)
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
