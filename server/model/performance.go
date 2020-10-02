package model

//Performance defines a connection between a movie, the characters and the people who acted the role
type Performance struct {
	UID       string   `json:"uid"`
	Film      []Movie  `json:"performance.film"`
	Character []Person `json:"performance.character"`
	Actor     []Person `json:"performance.actor"`
}
