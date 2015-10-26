package main

import (
    "time"
    "google.golang.org/appengine"
)

type Editor struct {
    Name string `json:"name" datastore:"Name"`
    Email string `json:"email" datastore:"Email"`
    IconURL string `json:"icon_url" datastore:"IconURL"`
    UpdatedAt time.Time `json:"updated_at" datastore:"UpdatedAt"`
    CreatedAt time.Time `json:"created_at" datastore:"CreatedAt"`
}

type Spot struct {
    SpotCode string `json:"spot_code" datastore:"SpotCode"`
    Revision int `json:"revision" datastore:"Revision"`
    SpotName string `json:"spot_name" datastore:"SpotName"`
    Body string `json:"body" datastore:"Body,noindex"`
    Photos []string `json:"photos" datastore:"Photos"`
    GeoInfo appengine.GeoPoint `json:"geo_info" datastore:"GeoInfo"`
    Phone string `json:"phone" datastore:"Phone"`
    UpdatedAt time.Time `json:"updated_at" datastore:"UpdatedAt"`
    CreatedAt time.Time `json:"created_at" datastore:"CreatedAt"`
}
