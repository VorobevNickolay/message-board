package main

type Message struct {
	ID     string `json:"id"`
	UserId string `json:"UserId"`
	Text   string `json:"Text"`
}

var messages = []Message{
	{ID: "1", UserId: "1", Text: "Meow"},
	{ID: "2", UserId: "2", Text: "I'm not happy:("},
	{ID: "3", UserId: "2", Text: "Where is my food?"},
}
