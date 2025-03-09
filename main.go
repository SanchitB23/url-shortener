package main

import (
	"log"
	"url-shortener/app"
)

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		log.Panic("Error starting the app:", err)
	}
}
