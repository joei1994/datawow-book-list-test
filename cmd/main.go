package main

import (
	application "datawow/book-list"
	"datawow/book-list/config"
)

func main() {
	config := config.LoadConfig()

	application.Start(config)
}
