package main

import (
	"time"

	"google.golang.org/appengine"
)

// User is the kind which stores data of users who comment and have favourite list
type User struct {
	Name       string    `json:"name" datastore:"Name"`
	Email      string    `json:"email" datastore:"Email"`
	IconURL    string    `json:"icon_url" datastore:"IconURL"`
	Favourites []string  `json:"favourites" datastore:"Favourites"` // list of SpotCode
	IsEditor   bool      `json:"is_editor" datastore:"IsEditor"`
	IsAdmin    bool      `json:"is_admin" datastore:"IsAdmin"`
	UpdatedAt  time.Time `json:"updated_at" datastore:"UpdatedAt"`
	CreatedAt  time.Time `json:"created_at" datastore:"CreatedAt"`
}

// Spot is the kind which stores sightseeing spot information
// Set SpotCode as KeyName
type Spot struct {
	SpotCode       string             `json:"spot_code" datastore:"SpotCode"`
	RevisionNumber int                `json:"revision" datastore:"Revision"` // increment on update
	Status         string             `json:"status" datastore:"Status"`     // 'status' or 'revision'
	SpotName       string             `json:"spot_name" datastore:"SpotName"`
	Body           string             `json:"body" datastore:"Body,noindex"`
	Photos         []string           `json:"photos" datastore:"Photos"`
	GeoInfo        appengine.GeoPoint `json:"geo_info" datastore:"GeoInfo"`
	Phone          string             `json:"phone" datastore:"Phone"`
	UpdatedAt      time.Time          `json:"updated_at" datastore:"UpdatedAt"`
	CreatedAt      time.Time          `json:"created_at" datastore:"CreatedAt"`
}
