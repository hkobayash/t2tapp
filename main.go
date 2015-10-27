package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func init() {
	http.Handle("/", goji.DefaultMux)

	goji.Get("/", indexHandler)
}

func render(v string, w io.Writer, data map[string]interface{}) {
	tmpl := template.Must(template.ParseFiles("views/layout.html", v))
	tmpl.Execute(w, data)
}

func indexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Name": "home",
	}
	render("views/index.html", w, data)
}
