package main

import (
	"net/http"

	"github.com/zenazn/goji"
)

func init() {
	http.Handle("/", goji.DefaultMux)

	goji.Get("/", indexHandler)
	goji.Get("/api/v1/spots", spotHandler)
	goji.Get("/edit/v1/spots", spotHandler)
	goji.Post("/edit/v1/spots", spotCreateHandler)
	//	goji.Patch("/edit/v1/spots/:spotCode", spotUpdateHandler)
	//	goji.Delete("/edit/v1/spots/:spotCode", spotDeleteHandler)
}
