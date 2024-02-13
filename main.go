package main

import (
	"fmt"
	"github.com/mkuokkanen/servicemap/configuration"
	"github.com/mkuokkanen/servicemap/model"
	"github.com/mkuokkanen/servicemap/server"
	"os"
)

func main() {
	fmt.Println("Starting application")
	// Config
	c := configuration.NewConfig()
	// Model
	m, err := model.NewModel(c)
	if err != nil {
		fmt.Printf("error building model: %s\n", err)
		os.Exit(1)
	}
	server.StartHttpServer(c, m)
}
