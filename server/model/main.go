package model

import "time"

type Movie struct {
	UID                 string        `json:"uid,omitempty"`
	Name                string        `json:"name@en"`
	Genre               []Genre       `json:"genre"`
	InitialRelreaseDate time.Time     `json:"initial_release_date"`
	Starring            []Performance `json:"starring"`
	Directors           []Person      `json:"~director.film"`
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
	UID           string        `json:"uid,omitempty"`
	FilmsActed    []Performance `json:"actor.film"`
	FilmsDirected []Movie       `json:"director.film"`
	Name          string        `json:"name@en"`
}

// Act flattens the object and returns Actor,
// whose type is easier to work with and easillly mapped on the fronend
// This functionality could have been easily achieved at the query level, but I
// Wanted to maintain 'purity' at that level hence easier for anyone to understand
// Implementing at db level would have also meant that I alter the model, which
// Would have made understanding 'relationships' harder since not every model
// would be directly mapped to the db-level types
func (p Person) Act() *Actor {
	return &Actor{
		UID:  p.UID,
		Name: p.Name,
		FilmsActed: func() []Film {
			movies := []Film{}
			for _, acted := range p.FilmsActed {
				for _, movie := range acted.Film {
					movies = append(movies, *movie.ConvertFilm())
				}
			}
			return movies
		}(),
	}
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

//FlatWithName is a struct with just UID and name fields
type FlatWithName struct {
	Name string `json:"name"`
	UID  string `json:"uid"`
}

// Film is a flattened form of movie with a direct connection to the actors
type Film struct {
	UID         string         `json:"uid,omitempty"`
	Name        string         `json:"name"`
	Genre       []string       `json:"genre"`
	ReleaseDate time.Time      `json:"release_date"`
	Actors      []FlatWithName `json:"actors"`
	Directors   []FlatWithName `json:"directors"`
}

// Actor is a flattened form of Person
//with a direct connection to the movie acted
type Actor struct {
	UID        string `json:"uid,omitempty"`
	FilmsActed []Film `json:"films_acted" `
	Name       string `json:"name"`
}
