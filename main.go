package main

import (
	"github.com/mkuokkanen/servicemap/configuration"
	"github.com/mkuokkanen/servicemap/model"
	"github.com/mkuokkanen/servicemap/server"
	"log"
	"os"
)

func main() {
	log.Printf("Starting application")
	// Config
	c := configuration.NewConfig()
	// Model
	m, err := model.NewModel(c)
	if err != nil {
		log.Printf("error building model: %s", err)
		os.Exit(1)
	}
	server.StartHttpServer(c, m)
}
