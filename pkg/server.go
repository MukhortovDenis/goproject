package pkg

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(path string, router http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    path,
		Handler: router,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
