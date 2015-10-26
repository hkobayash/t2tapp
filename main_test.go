package main

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

