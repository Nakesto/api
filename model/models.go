package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Pswd string `json:"pswd"`
}
