package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var tmpl *template.Template

type Address struct {
	addr string
}

func main() {
	tmpl = template.Must(template.ParseFiles("/var/www/html/index.html"))
	//tmpl = template.Must(template.ParseFiles("web/index.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
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
