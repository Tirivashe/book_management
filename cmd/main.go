package main

import (
	"log"

	"github.com/Tirivashe/book_management/config"
)

func main() {
	app := config.NewAPIServer(":3000")
	log.Fatal(app.Start())
}