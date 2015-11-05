package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"

	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/user"
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
		t.Fatalf("Failed to create req: %v", err)
	}
	loginUser := user.User{Email: "hoge@gmail.com", Admin: false, ID: "111111"}
	aetest.Login(&loginUser, req)
	_ = appengine.NewContext(req)
	res := httptest.NewRecorder()
	c := web.C{}
	spotHandler(c, res, req)
	if res.Code != http.StatusOK {
		t.Fatalf("Fail to request spots list")
	}
}

func TestCreateSpot(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	input, err := json.Marshal(Spot{SpotName: "foo", Body: "bar"})
	req, err := inst.NewRequest("POST", "/edit/v1/spots", bytes.NewBuffer(input))
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	loginUser := user.User{Email: "hoge@gmail.com", Admin: false, ID: "111111"}
	aetest.Login(&loginUser, req)
	_ = appengine.NewContext(req)
	res := httptest.NewRecorder()
	c := web.C{}
	spotCreateHandler(c, res, req)
	if res.Code != http.StatusCreated {
		t.Fatalf("Fail to request spots create, status code: %v", res.Code)
	}
}
