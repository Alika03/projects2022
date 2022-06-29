package main

import (
	"back-end/server"
	"log"
)

func main() {
	app := server.NewApp()

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
