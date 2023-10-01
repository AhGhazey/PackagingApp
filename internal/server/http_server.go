package server

import (
	"context"
	"fmt"
	"github/ahmedghazey/packaging/internal/configuration"
	"net/http"
)

type HttpServer struct {
	server  *http.Server
	address string
}

func NewHttpServer(router http.Handler) *HttpServer {
	config, _ := configuration.GetConfiguration()
	server := &http.Server{
		Addr:         config.ServerAddress,
		Handler:      router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}
	return &HttpServer{
		server:  server,
		address: config.ServerAddress,
	}
}

// Run starts the HTTP server.
func (s *HttpServer) Run() error {
	fmt.Printf("Server listening on %s", s.address)
	return s.server.ListenAndServe()
}

// Stop gracefully stops the HTTP server.
func (s *HttpServer) Stop(ctx context.Context) error {
	fmt.Println("Stopping server...")
	return s.server.Shutdown(ctx)
}
