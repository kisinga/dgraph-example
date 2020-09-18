package model

import "time"

type Movie struct {
	UID                 string        `json:"uid,omitempty"`
	Name                string        `json:"name@en"`
	Genre               []Genre       `json:"genre"`
	InitialRelreaseDate time.Time     `json:"initial_release_date"`
	Starring            []Performance `json:"starring" dgraph:"count reverse"`
}

type Performance struct {
	UID       string   `json:"uid"`
	Film      []Movie  `json:"performance.film"`
	Character []Person `json:"performance.character"`
	Actor     []Person `json:"performance.actor"`
}

type Genre struct {
	Name string `json:"name@en"`
}
type Person struct {
	FilmsActed    []Performance `json:"actor.film" dgraph:"count reverse"`
	FilmsDirected []Movie       `json:"director.film" dgraph:"count reverse"`
	Name          string        `json:"name@en" dgraph:"count reverse"`
}
