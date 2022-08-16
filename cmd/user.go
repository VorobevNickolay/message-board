package main

type User struct {
	ID       string `json:"id"`
	Username string `json:"Username" gorm:"unique"`
}

var users = []User{
	{ID: "1", Username: "Garfield"},
	{ID: "2", Username: "Pirate"},
}
