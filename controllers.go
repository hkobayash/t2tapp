package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
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

func spotGetHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ren := render.New()
	spotCode, err := strconv.ParseInt(c.URLParams["spotCode"], 10, 64)
	if err != nil {
		log.Infof(ctx, "failed to parse int :%v", err)
		ren.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "error, bad resource id"})
		return
	}
	var spot Spot
	spot.SpotCode = spotCode
	err = datastore.Get(ctx, spot.key(ctx), &spot)
	if err != nil {
		ren.JSON(w, http.StatusNotFound, map[string]interface{}{"message": "error, entity not found"})
		return
	}
	ren.JSON(w, http.StatusOK, map[string]interface{}{"item": spot})
}

func spotCreateHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	var spot Spot
	ctx := appengine.NewContext(r)
	ren := render.New()
	err := json.NewDecoder(r.Body).Decode(&spot)
	if err != nil {
		ren.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "error, cannot decode json"})
		log.Infof(ctx, "failed to parse JSON:%v", err)
		return
	}
	_, err = spot.Create(ctx)
	if err != nil {
		ren.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "error, cannot create new entity"})
		return
	}
	ren.JSON(w, http.StatusCreated, map[string]interface{}{"message": "new entity created", "item": spot})
}

func spotUpdateHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ren := render.New()
	spotCode, err := strconv.ParseInt(c.URLParams["spotCode"], 10, 64)
	if err != nil {
		ren.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "error, bad resource id"})
		return
	}
	var checkSpot Spot
	checkSpot.SpotCode = spotCode
	currentUser := user.Current(ctx)
	err = datastore.Get(ctx, checkSpot.key(ctx), &checkSpot)
	if err != nil {
		ren.JSON(w, http.StatusNotFound, map[string]interface{}{"message": "error, entity not found"})
		return
	}
	if currentUser.ID != checkSpot.Editor {
		ren.JSON(w, http.StatusForbidden, map[string]interface{}{"message": "error, you don't have right to write this entity"})
		return
	}
	var updateSpot Spot
	err = json.NewDecoder(r.Body).Decode(&updateSpot)
	if err != nil {
		ren.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "error, cannot decode json"})
		log.Infof(ctx, "failed to parse JSON:%v", err)
		return
	}
	updateSpot.Editor = checkSpot.Editor
	_, err = updateSpot.Update(ctx, spotCode)
	if err != nil {
		ren.JSON(w, http.StatusForbidden, map[string]interface{}{"message": "error, you cannot edit this spot"})
		return
	}
	ren.JSON(w, http.StatusCreated, map[string]interface{}{"message": "entity updated", "item":updateSpot})
}
