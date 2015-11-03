package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestIndex(t *testing.T) {
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

func TestSpot(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/edit/v1/spots", nil)
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	_ = appengine.NewContext(req)
	res := httptest.NewRecorder()
	c := web.C{}
	spotHandler(c, res, req)
}

func TestCreateSpot(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	input, err := json.Marshal(Spot{SpotName: "foo", Body: "bar"})
	req, err := inst.NewRequest("POST", "/edit/v1/spots", bytes.NewBuffer(input))
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	_ = appengine.NewContext(req)
}
