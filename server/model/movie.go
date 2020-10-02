package model

import "time"


//Movie is a struct of 
type Movie struct {
	UID                 string        `json:"uid,omitempty"`
	Name                string        `json:"name@en"`
	Genre               []Genre       `json:"genre"`
	InitialRelreaseDate time.Time     `json:"initial_release_date"`
	Starring            []Performance `json:"starring"`
	Directors           []Person      `json:"~director.film"`
}

//ConvertFilm converts a movie to the film equivalent for easy mapping
//See documentation for act for further understanding
func (m Movie) ConvertFilm() *Film {
	return &Film{
		UID:         m.UID,
		Name:        m.Name,
		ReleaseDate: m.InitialRelreaseDate,
		Genre: func() []string {
			genres := []string{}
			for _, genre := range m.Genre {
				genres = append(genres, genre.Name)
			}
			return genres
		}(),
		Actors: func() []FlatWithName {
			actors := []FlatWithName{}
			for _, starring := range m.Starring {
				for _, actor := range starring.Actor {
					actors = append(actors, FlatWithName{
						Name: actor.Name,
						UID:  actor.UID,
					})
				}
			}
			return actors
		}(),
		Directors: func() []FlatWithName {
			directors := []FlatWithName{}
			for _, director := range m.Directors {
				directors = append(directors, FlatWithName{
					Name: director.Name,
					UID:  director.UID,
				})
			}
			return directors
		}(),
	}
}
