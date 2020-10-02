package model

import "time"

// Film is a flattened form of movie with a direct connection to the actors
type Film struct {
	UID         string         `json:"uid,omitempty"`
	Name        string         `json:"name"`
	Genre       []string       `json:"genre"`
	ReleaseDate time.Time      `json:"release_date"`
	Actors      []FlatWithName `json:"actors"`
	Directors   []FlatWithName `json:"directors"`
}
