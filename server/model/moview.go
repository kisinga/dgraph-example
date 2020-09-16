package model

import (
	"time"
)

type Movie struct {
	UID                 string      `json:"uid,omitempty"`
	Name                string      `json:"name@en"`
	Starring            string      `json:"starring"`
	Genre               string      `json:"genre"`
	InitialRelreaseDate time.Time   `json:"initial_relrease_date"`
	Performance         Performance `json:"performance"`
}

type Performance struct {
	actor Person
}
