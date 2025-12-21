package server

import (
	"net/http"

	"advanced-go-app/internal/logger"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	logger.Logger.Println("Server Started on", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
