package main

import (
	"back-end/config"
	"back-end/server"
	"back-end/utils"
	"log"
)

func init() {
	if err := utils.CreateKey("privateKey.pem", "publicKey.pem"); err != nil {
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
