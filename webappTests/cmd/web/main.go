package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	// set up an app config
	app := application{}

	// get application routes
	mux := app.routes()

	// print out a message to say the server is running
	log.Println("Starting server on :8080")

	// start the web server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
