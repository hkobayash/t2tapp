package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

/*
import (
    "testing"
    "google.golang.org/appengine/aetest"
)

func TestIndexHandler(t *testing.T) {
    inst, err := aetest.NewInstance(nil)
    if err != nil {
        t.Fatalf("Failed to create instance: %v", err)
    }
    defer inst.Close()

    _, err = inst.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatalf("Failed to create req1: %v", err)
    }
}
*/

func Route(m *web.Mux) {
	m.Get("/", indexHandler)
	m.Get("/edit/v1/spots", spotHandler)
	m.Post("/edit/v1/spots", spotCreateHandler)
}

func TestIndex(t *testing.T) {
	m := web.New()
	Route(m)
	ts := httptest.NewServer(m)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Error("unexpected")
	}
	if res.StatusCode != http.StatusOK {
		t.Error("invalid status code")
	}
	res2, err := http.Get(ts.URL + "/hoge")
	if err != nil {
		t.Error("unexpected")
	}
	if res2.StatusCode != http.StatusNotFound {
		t.Error("invalid status code")
	}
}

func TestSpot(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req1, err := inst.NewRequest("GET", "/edit/v1/spots", nil)
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	_ = appengine.NewContext(req1)

	input, err := json.Marshal(Spot{SpotName: "foo", Body: "bar"})
	req2, err := inst.NewRequest("POST", "/edit/v1/spots", bytes.NewBuffer(input))
	if err != nil {
		t.Fatalf("Failed to create req2: %v", err)
	}
	_ = appengine.NewContext(req2)
	/*
	input, err := json.Marshal(Spot{SpotName: "foo", Body: "bar"})
	res_post, err := http.Post(ts.URL+"/edit/v1/spots", "application/json", bytes.NewBuffer(input))
	if err != nil {
		t.Error("unexpected")
	}
	if res_post.StatusCode != http.StatusCreated {
		t.Error("invalid status code")
	}*/
}
