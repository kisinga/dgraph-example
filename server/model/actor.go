package model

// Actor is a flattened form of Person
//with a direct connection to the movie acted
type Actor struct {
	UID        string `json:"uid,omitempty"`
	FilmsActed []Film `json:"films_acted" `
	Name       string `json:"name"`
}
