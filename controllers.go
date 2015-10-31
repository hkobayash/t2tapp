package main

import (
	"net/http"

	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func indexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	ren := render.New(render.Options{
		Layout:     "layout",
		Extensions: []string{".html"},
	})
	ren.HTML(w, http.StatusOK, "index", nil)
}

func spotHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	ren := render.New()
	con := appengine.NewContext(r)
	spots := []Spot{}
	_, err := datastore.NewQuery("Spot").Order("-UpdatedAt").GetAll(con, &spots)
	if err != nil {
		ren.JSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error"})
	}
	ren.JSON(w, http.StatusOK, map[string]interface{}{"items": spots})
}
