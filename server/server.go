package server

import (
	"errors"
	"fmt"
	"github.com/mkuokkanen/servicemap/configuration"
	"github.com/mkuokkanen/servicemap/model"
	"net/http"
	"os"
)

type Server struct {
	config *configuration.Config
	model  *model.Model
}

func StartHttpServer(c *configuration.Config, m *model.Model) {
	// Setup Server
	server := &Server{
		config: c,
		model:  m,
	}
	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", server.rootPath)
	mux.HandleFunc("GET /index.html", server.renderIndex)
	mux.HandleFunc("POST /reload", server.reloadData)
	mux.Handle("GET /static/", pathStatic(c.StaticDir))
	// server
	err := http.ListenAndServe(":"+c.HttpPort, mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
