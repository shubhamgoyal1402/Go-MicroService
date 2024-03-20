package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = 80

type Config struct {
}

func main() {

	app := Config{}

	log.Printf("Starting Broker Service on port %v\n", webPort)

	// Define http server

	srv := &http.Server{

		Addr:    fmt.Sprintf(":%v", webPort),
		Handler: app.routes(),
	}

	//Start the Server
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
