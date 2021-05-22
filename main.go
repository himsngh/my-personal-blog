package main

import (
	"github.com/himsngh/my-personal-blog/routes"
	"log"
	"net/http"
)

func app() error {

	handler := routes.Initialize()
	server := &http.Server{
		Addr: ":8080",
		Handler: handler,
	}

	return server.ListenAndServe()
}

func main() {

	if err := app(); err != nil {
		log.Fatal("error running the server " + err.Error())
	}
}