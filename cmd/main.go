package main

import "message-board/Internal/app"

func main() {
	router := app.NewRouter()
	router.SetUpRouter()
}
