package model

//Person defined a human that either directed or acted a movie
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
