package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/zenazn/goji/web"
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
