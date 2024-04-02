package server

import (
	"errors"
	"github.com/mkuokkanen/servicemap/configuration"
	"github.com/mkuokkanen/servicemap/model"
	"log"
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
	log.Printf("Starting server at port %s", c.HttpPort)
	logMux := LogMiddleware(mux)
	err := http.ListenAndServe(":"+c.HttpPort, logMux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed")
	} else if err != nil {
		log.Printf("error starting server: %s", err)
		os.Exit(1)
	}
}
