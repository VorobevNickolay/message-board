package main

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

var users = []User{
	{ID: 1, Username: "Garfield", Password: "Orange123"},
	{ID: 2, Username: "Pirate", Password: "QuartoChampion"},
}
