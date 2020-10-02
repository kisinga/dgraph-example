package model

//FlatWithName is a struct with just UID and name fields
type FlatWithName struct {
	Name string `json:"name"`
	UID  string `json:"uid"`
}
