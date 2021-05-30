package main

import (
	"github.com/himsngh/my-personal-blog/server"
	"log"
	"net/http"
)

func app() error {

	routesServer, err := server.NewServer()
	if err != nil {
		return err
	}

	handler, err := routesServer.ServeRoutes()
	if err != nil {
		return err
	}

	serve := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	return serve.ListenAndServe()
}

func main() {

	if err := app(); err != nil {
		log.Fatal("error running the server " + err.Error())
	}
}
