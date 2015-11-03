package main

import (
	"encoding/json"
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
	ctx := appengine.NewContext(r)
	spots := []Spot{}
	_, err := datastore.NewQuery("Spot").Order("-UpdatedAt").GetAll(ctx, &spots)
	if err != nil {
		ren.JSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error"})
		return
	}
	ren.JSON(w, http.StatusOK, map[string]interface{}{"items": spots})
}

func spotCreateHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	var spot Spot
	ctx := appengine.NewContext(r)
	ren := render.New()
	err := json.NewDecoder(r.Body).Decode(spot)
	if err != nil {
		ren.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "error"})
		return
	}
	_, err = spot.Create(ctx)
	if err != nil {
		ren.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "error"})
		return
	}
	ren.JSON(w, http.StatusCreated, map[string]interface{}{"message": "new entity created"})
}
