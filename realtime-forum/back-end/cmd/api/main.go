package main

import (
	"back-end/config"
	"back-end/pkg"
	"back-end/server"
	"log"
)

func init() {
	if err := pkg.CreateKey("privateKey.pem", "publicKey.pem"); err != nil {
		panic(err)
	}
	config.GetConfig()
}

func main() {
	app := server.NewApp()

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
