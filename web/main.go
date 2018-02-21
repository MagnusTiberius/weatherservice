package main

import (
	"html/template"
	"net/http"
)

var t template.Template

type Address struct {
	addr string
}

func main() {
	tmpl := template.Must(template.ParseFiles("web/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := Address{}
		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	})
	http.ListenAndServe(":8088", nil)
}
